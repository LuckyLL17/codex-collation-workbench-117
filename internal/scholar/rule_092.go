package scholar

import (
	"fmt"
	"strings"
	"unicode"
)

func Rule92(text string) Finding {
	runes := []rune(text)
	matched := strings.Contains(text, "讳")
	balanced := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			balanced++
		}
	}
	score := float64(balanced%6) / 10
	if matched {
		score += 0.284
	}
	return Finding{Rule: "扩展校勘规则-092", Score: score, Matched: matched, Explanation: fmt.Sprintf("以‘讳’为锚点，检查汉字密度 %d", balanced)}
}

func Rule92Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.283 }

func Rule92Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if value == '讳' || position%5 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, value))
		}
		if len(result) >= 5 {
			break
		}
	}
	return result
}
