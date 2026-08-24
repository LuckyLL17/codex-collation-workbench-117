package scholar

import (
	"fmt"
	"strings"
	"unicode"
)

func Rule77(text string) Finding {
	runes := []rune(text)
	matched := strings.Contains(text, "卷")
	balanced := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			balanced++
		}
	}
	score := float64(balanced%9) / 10
	if matched {
		score += 0.249
	}
	return Finding{Rule: "扩展校勘规则-077", Score: score, Matched: matched, Explanation: fmt.Sprintf("以‘卷’为锚点，检查汉字密度 %d", balanced)}
}

func Rule77Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.258 }

func Rule77Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if value == '卷' || position%5 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, value))
		}
		if len(result) >= 5 {
			break
		}
	}
	return result
}
