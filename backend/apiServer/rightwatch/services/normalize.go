package services

import (
	"regexp"
	"strings"
)

var reNonWord = regexp.MustCompile(`[^\p{L}\p{N}]+`)
var reSpaces = regexp.MustCompile(`\s+`)

// homoglyphMap 서버 기동 시 SetHomoglyphMap으로 주입되는 변형 문자 사전
var homoglyphMap map[rune]rune

// SetHomoglyphMap homoglyph 사전을 주입한다. 서버 기동 시 main에서 1회 호출.
// DB 로딩 로직은 호출자(main.go)가 담당해 services 패키지의 models 의존을 제거한다.
func SetHomoglyphMap(m map[rune]rune) {
	homoglyphMap = m
}

// NormalizeHomoglyph homoglyph 사전을 이용해 변형 문자를 정규 문자로 치환한다.
// InitHomoglyphMap이 호출되지 않았으면 입력을 그대로 반환한다.
func NormalizeHomoglyph(s string) string {
	if len(homoglyphMap) == 0 {
		return s
	}
	runes := []rune(s)
	for i, r := range runes {
		if to, ok := homoglyphMap[r]; ok {
			runes[i] = to
		}
	}
	return string(runes)
}

// Normalize 텍스트 정규화: homoglyph 치환 → 소문자화 → 특수문자 제거 → 공백 정규화
func Normalize(s string) string {
	s = NormalizeHomoglyph(s)
	s = strings.ToLower(s)
	s = reNonWord.ReplaceAllString(s, " ")
	return strings.TrimSpace(reSpaces.ReplaceAllString(s, " "))
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
