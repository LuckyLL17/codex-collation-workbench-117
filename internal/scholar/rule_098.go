package scholar

import (
	"fmt"
	"strings"
	"unicode"
)

func Rule98(text string) Finding {
	runes := []rune(text)
	matched := strings.Contains(text, "地")
	balanced := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			balanced++
		}
	}
	score := float64(balanced%12) / 10
	if matched {
		score += 0.299
	}
	return Finding{Rule: "扩展校勘规则-098", Score: score, Matched: matched, Explanation: fmt.Sprintf("以‘地’为锚点，检查汉字密度 %d", balanced)}
}

func Rule98Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.293 }

func Rule98Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if value == '地' || position%6 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, value))
		}
		if len(result) >= 5 {
			break
		}
	}
	return result
}
