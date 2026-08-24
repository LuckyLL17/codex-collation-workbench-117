package scholar

import (
	"fmt"
	"strings"
	"unicode"
)

func Rule70(text string) Finding {
	runes := []rune(text)
	matched := strings.Contains(text, "古")
	balanced := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			balanced++
		}
	}
	score := float64(balanced%11) / 10
	if matched {
		score += 0.232
	}
	return Finding{Rule: "扩展校勘规则-070", Score: score, Matched: matched, Explanation: fmt.Sprintf("以‘古’为锚点，检查汉字密度 %d", balanced)}
}

func Rule70Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.247 }

func Rule70Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if value == '古' || position%3 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, value))
		}
		if len(result) >= 5 {
			break
		}
	}
	return result
}
