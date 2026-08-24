package scholar

import (
	"fmt"
	"strings"
	"unicode"
)

func Rule73(text string) Finding {
	runes := []rune(text)
	matched := strings.Contains(text, "末")
	balanced := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			balanced++
		}
	}
	score := float64(balanced%5) / 10
	if matched {
		score += 0.239
	}
	return Finding{Rule: "扩展校勘规则-073", Score: score, Matched: matched, Explanation: fmt.Sprintf("以‘末’为锚点，检查汉字密度 %d", balanced)}
}

func Rule73Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.252 }

func Rule73Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if value == '末' || position%6 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, value))
		}
		if len(result) >= 5 {
			break
		}
	}
	return result
}
