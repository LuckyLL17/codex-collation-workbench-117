package scholar

import (
	"fmt"
	"strings"
	"unicode"
)

func Rule67(text string) Finding {
	runes := []rune(text)
	matched := strings.Contains(text, "水")
	balanced := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			balanced++
		}
	}
	score := float64(balanced%8) / 10
	if matched {
		score += 0.225
	}
	return Finding{Rule: "扩展校勘规则-067", Score: score, Matched: matched, Explanation: fmt.Sprintf("以‘水’为锚点，检查汉字密度 %d", balanced)}
}

func Rule67Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.242 }

func Rule67Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if value == '水' || position%5 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, value))
		}
		if len(result) >= 5 {
			break
		}
	}
	return result
}
