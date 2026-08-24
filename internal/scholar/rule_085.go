package scholar

import (
	"fmt"
	"strings"
	"unicode"
)

func Rule85(text string) Finding {
	runes := []rune(text)
	matched := strings.Contains(text, "同")
	balanced := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			balanced++
		}
	}
	score := float64(balanced%8) / 10
	if matched {
		score += 0.268
	}
	return Finding{Rule: "扩展校勘规则-085", Score: score, Matched: matched, Explanation: fmt.Sprintf("以‘同’为锚点，检查汉字密度 %d", balanced)}
}

func Rule85Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.272 }

func Rule85Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if value == '同' || position%3 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, value))
		}
		if len(result) >= 5 {
			break
		}
	}
	return result
}
