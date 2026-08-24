package lexicon

import "strings"

func Similarity(a, b string) float64 {
	a = Normalize(a)
	b = Normalize(b)
	if a == b && a != "" {
		return 1
	}
	if a == "" || b == "" {
		return 0
	}
	common := 0
	for _, r := range a {
		if strings.ContainsRune(b, r) {
			common++
		}
	}
	return float64(common*2) / float64(len([]rune(a))+len([]rune(b)))
}
func Tokens(value string) []string { return strings.Fields(Normalize(value)) }
