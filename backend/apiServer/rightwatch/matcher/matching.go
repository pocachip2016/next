package matcher

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"rightwatch/models"
	"rightwatch/services"
)

// SynonymMap maps each normalized word to all words in its synonym group (including itself)
type SynonymMap map[string][]string

// MatchResult holds one matched (post, content) pair
type MatchResult struct {
	ContentID uint
	PostID    string
	PostIdx   string
	PostTxt   string
}

// loadSynonymMap loads synonym_words from DB and groups by pair_id
func loadSynonymMap(db *gorm.DB) (SynonymMap, error) {
	var words []models.SynonymWord
	if err := db.Find(&words).Error; err != nil {
		return nil, err
	}

	groups := make(map[int][]string)
	for _, w := range words {
		if n := services.Normalize(w.Synonym); n != "" {
			groups[w.PairId] = append(groups[w.PairId], n)
		}
	}

	sm := make(SynonymMap)
	for _, syns := range groups {
		for _, s := range syns {
			sm[s] = syns
		}
	}
	return sm, nil
}

// ExpandWithSynonyms returns candidate token sets: each token is replaced with its synonym group
func ExpandWithSynonyms(tokens []string, sm SynonymMap) [][]string {
	expanded := make([][]string, len(tokens))
	for i, tok := range tokens {
		if syns, ok := sm[tok]; ok {
			expanded[i] = syns
		} else {
			expanded[i] = []string{tok}
		}
	}
	return expanded
}

// matchExpanded returns true if text contains at least one candidate from every token group
func matchExpanded(text string, expanded [][]string) bool {
	for _, candidates := range expanded {
		found := false
		for _, c := range candidates {
			if strings.Contains(text, c) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// RunMatching matches all posts against all kta_contents using normalized title + synonyms.
// Returns new matches not already present in check_list.
// DB writes are left to the caller (handler).
func RunMatching(db *gorm.DB) ([]MatchResult, error) {
	sm, err := loadSynonymMap(db)
	if err != nil {
		return nil, fmt.Errorf("loadSynonymMap: %w", err)
	}

	var contents []models.KtaContent
	if err := db.Find(&contents).Error; err != nil {
		return nil, fmt.Errorf("load kta_contents: %w", err)
	}

	var posts []models.Post
	if err := db.Find(&posts).Error; err != nil {
		return nil, fmt.Errorf("load posts: %w", err)
	}

	var existing []models.CheckList
	if err := db.Find(&existing).Error; err != nil {
		return nil, fmt.Errorf("load check_list: %w", err)
	}
	seen := make(map[string]bool, len(existing))
	for _, cl := range existing {
		seen[fmt.Sprintf("%d:%s", cl.ContentId, cl.PostIdx)] = true
	}

	var results []MatchResult
	for _, content := range contents {
		tokens := services.Tokens(content.Title)
		if len(tokens) == 0 {
			continue
		}
		expanded := ExpandWithSynonyms(tokens, sm)

		for _, post := range posts {
			postNorm := services.Normalize(post.Txt)
			if !matchExpanded(postNorm, expanded) {
				continue
			}
			key := fmt.Sprintf("%d:%s", content.Id, post.Idx)
			if seen[key] {
				continue
			}
			seen[key] = true
			results = append(results, MatchResult{
				ContentID: content.Id,
				PostID:    fmt.Sprintf("%d", post.Id),
				PostIdx:   post.Idx,
				PostTxt:   post.Txt,
			})
		}
	}
	return results, nil
}
