package scholar

import (
	"fmt"
	"strings"
	"unicode"
)

func Rule99(text string) Finding {
	runes := []rune(text)
	matched := strings.Contains(text, "人")
	balanced := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			balanced++
		}
	}
	score := float64(balanced%4) / 10
	if matched {
		score += 0.301
	}
	return Finding{Rule: "扩展校勘规则-099", Score: score, Matched: matched, Explanation: fmt.Sprintf("以‘人’为锚点，检查汉字密度 %d", balanced)}
}

func Rule99Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.295 }

func Rule99Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if value == '人' || position%7 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, value))
		}
		if len(result) >= 5 {
			break
		}
	}
	return result
}
