package services

import (
	"regexp"
	"strings"
)

var reNonWord = regexp.MustCompile(`[^\p{L}\p{N}]+`)

// Normalize 텍스트 정규화
// 특수문자·공백 제거, 소문자화, 공백 정규화
func Normalize(s string) string {
	s = strings.ToLower(s)
	s = reNonWord.ReplaceAllString(s, " ")
	return strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(s, " "))
}

// Tokens 정규화 후 공백으로 토큰 분리
func Tokens(s string) []string {
	n := Normalize(s)
	if n == "" {
		return nil
	}
	return strings.Fields(n)
}

// ContainsAll 대상 문자열이 모든 토큰을 포함하는지 검사
func ContainsAll(target string, tokens []string) bool {
	t := Normalize(target)
	for _, tok := range tokens {
		if !strings.Contains(t, tok) {
			return false
		}
	}
	return true
}
