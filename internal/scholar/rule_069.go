package scholar

import (
	"fmt"
	"strings"
	"unicode"
)

func Rule69(text string) Finding {
	runes := []rune(text)
	matched := strings.Contains(text, "月")
	balanced := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			balanced++
		}
	}
	score := float64(balanced%10) / 10
	if matched {
		score += 0.230
	}
	return Finding{Rule: "扩展校勘规则-069", Score: score, Matched: matched, Explanation: fmt.Sprintf("以‘月’为锚点，检查汉字密度 %d", balanced)}
}

func Rule69Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.245 }

func Rule69Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if value == '月' || position%7 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, value))
		}
		if len(result) >= 5 {
			break
		}
	}
	return result
}
