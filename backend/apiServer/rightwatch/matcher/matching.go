package matcher

import (
	"fmt"
	"strings"
	"time"

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

// runMatchingWithPosts is the core matching loop shared by RunMatching, RunMatchingSince, RunMatchingForContent.
func runMatchingWithPosts(db *gorm.DB, contents []models.KtaContent, posts []models.Post) ([]MatchResult, error) {
	sm, err := loadSynonymMap(db)
	if err != nil {
		return nil, fmt.Errorf("loadSynonymMap: %w", err)
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

// RunMatching matches all posts against all kta_contents.
func RunMatching(db *gorm.DB) ([]MatchResult, error) {
	var contents []models.KtaContent
	if err := db.Find(&contents).Error; err != nil {
		return nil, fmt.Errorf("load kta_contents: %w", err)
	}
	var posts []models.Post
	if err := db.Find(&posts).Error; err != nil {
		return nil, fmt.Errorf("load posts: %w", err)
	}
	return runMatchingWithPosts(db, contents, posts)
}

// RunMatchingSince matches only posts updated on or after since against all kta_contents.
// If since is zero, falls back to a full scan (same as RunMatching).
func RunMatchingSince(db *gorm.DB, since time.Time) ([]MatchResult, error) {
	var contents []models.KtaContent
	if err := db.Find(&contents).Error; err != nil {
		return nil, fmt.Errorf("load kta_contents: %w", err)
	}
	var posts []models.Post
	q := db
	if !since.IsZero() {
		q = db.Where("last_update >= ?", since.Format("2006-01-02 15:04:05"))
	}
	if err := q.Find(&posts).Error; err != nil {
		return nil, fmt.Errorf("load posts: %w", err)
	}
	return runMatchingWithPosts(db, contents, posts)
}

// RunMatchingForContent matches all posts against a single kta_content (full scan).
// Used when a new content is registered and needs immediate detection.
func RunMatchingForContent(db *gorm.DB, contentID uint) ([]MatchResult, error) {
	var content models.KtaContent
	if db.First(&content, contentID).RecordNotFound() {
		return nil, fmt.Errorf("content %d not found", contentID)
	}
	var posts []models.Post
	if err := db.Find(&posts).Error; err != nil {
		return nil, fmt.Errorf("load posts: %w", err)
	}
	return runMatchingWithPosts(db, []models.KtaContent{content}, posts)
}
