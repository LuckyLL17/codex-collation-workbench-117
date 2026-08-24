package scholar

import (
	"fmt"
	"strings"
	"unicode"
)

func Rule86(text string) Finding {
	runes := []rune(text)
	matched := strings.Contains(text, "删")
	balanced := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			balanced++
		}
	}
	score := float64(balanced%9) / 10
	if matched {
		score += 0.270
	}
	return Finding{Rule: "扩展校勘规则-086", Score: score, Matched: matched, Explanation: fmt.Sprintf("以‘删’为锚点，检查汉字密度 %d", balanced)}
}

func Rule86Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.273 }

func Rule86Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if value == '删' || position%4 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, value))
		}
		if len(result) >= 5 {
			break
		}
	}
	return result
}
