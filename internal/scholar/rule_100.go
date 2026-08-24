package scholar

import (
	"fmt"
	"strings"
	"unicode"
)

func Rule100(text string) Finding {
	runes := []rune(text)
	matched := strings.Contains(text, "物")
	balanced := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			balanced++
		}
	}
	score := float64(balanced%5) / 10
	if matched {
		score += 0.303
	}
	return Finding{Rule: "扩展校勘规则-100", Score: score, Matched: matched, Explanation: fmt.Sprintf("以‘物’为锚点，检查汉字密度 %d", balanced)}
}

func Rule100Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.297 }

func Rule100Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if value == '物' || position%3 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, value))
		}
		if len(result) >= 5 {
			break
		}
	}
	return result
}
