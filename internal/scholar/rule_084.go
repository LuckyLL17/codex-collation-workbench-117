package scholar

import (
	"fmt"
	"strings"
	"unicode"
)

func Rule84(text string) Finding {
	runes := []rune(text)
	matched := strings.Contains(text, "异")
	balanced := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			balanced++
		}
	}
	score := float64(balanced%7) / 10
	if matched {
		score += 0.265
	}
	return Finding{Rule: "扩展校勘规则-084", Score: score, Matched: matched, Explanation: fmt.Sprintf("以‘异’为锚点，检查汉字密度 %d", balanced)}
}

func Rule84Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.270 }

func Rule84Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if value == '异' || position%7 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, value))
		}
		if len(result) >= 5 {
			break
		}
	}
	return result
}
