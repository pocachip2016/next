package services

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"드라마", "드라마"},
		{"Hello World", "hello world"},
		{"해피엔딩!!", "해피엔딩"},
		{"  spaces  ", "spaces"},
		{"A--B", "a b"},
	}
	for _, c := range cases {
		got := Normalize(c.input)
		if got != c.want {
			t.Errorf("Normalize(%q) = %q, want %q", c.input, got, c.want)
		}
	}
}

func TestNormalizeHomoglyph(t *testing.T) {
	SetHomoglyphMap(map[rune]rune{
		'0': 'o',
		'1': 'l',
		'3': 'e',
		'ㅇ': 'o',
	})
	defer func() { SetHomoglyphMap(nil) }()

	cases := []struct {
		input string
		want  string
	}{
		{"0ne", "one"},
		{"l33t", "leet"},
		{"ㅇ리온", "o리온"},
		{"정상텍스트", "정상텍스트"},
	}
	for _, c := range cases {
		got := NormalizeHomoglyph(c.input)
		if got != c.want {
			t.Errorf("NormalizeHomoglyph(%q) = %q, want %q", c.input, got, c.want)
		}
	}
}

func TestNormalizeWithHomoglyph(t *testing.T) {
	SetHomoglyphMap(map[rune]rune{
		'0': 'o',
		'1': 'l',
	})
	defer func() { SetHomoglyphMap(nil) }()

	cases := []struct {
		input string
		want  string
	}{
		{"h0ll0 w0rld", "hollo world"},
		{"1337", "l337"},
		{"정상 텍스트", "정상 텍스트"},
	}
	for _, c := range cases {
		got := Normalize(c.input)
		if got != c.want {
			t.Errorf("Normalize(%q) = %q, want %q", c.input, got, c.want)
		}
	}
}

func TestTokens(t *testing.T) {
	SetHomoglyphMap(map[rune]rune{'0': 'o'})
	defer func() { SetHomoglyphMap(nil) }()

	got := Tokens("드라마 0리온 100%")
	want := []string{"드라마", "o리온", "1oo"}
	if len(got) != len(want) {
		t.Fatalf("Tokens len=%d, want %d: %v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Tokens[%d]=%q, want %q", i, got[i], want[i])
		}
	}
}
