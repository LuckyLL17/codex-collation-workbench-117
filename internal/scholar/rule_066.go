package scholar

import (
	"fmt"
	"strings"
	"unicode"
)

func Rule66(text string) Finding {
	runes := []rune(text)
	matched := strings.Contains(text, "山")
	balanced := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			balanced++
		}
	}
	score := float64(balanced%7) / 10
	if matched {
		score += 0.222
	}
	return Finding{Rule: "扩展校勘规则-066", Score: score, Matched: matched, Explanation: fmt.Sprintf("以‘山’为锚点，检查汉字密度 %d", balanced)}
}

func Rule66Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.240 }

func Rule66Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if value == '山' || position%4 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, value))
		}
		if len(result) >= 5 {
			break
		}
	}
	return result
}
