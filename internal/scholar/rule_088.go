package scholar

import (
	"fmt"
	"strings"
	"unicode"
)

func Rule88(text string) Finding {
	runes := []rune(text)
	matched := strings.Contains(text, "脱")
	balanced := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			balanced++
		}
	}
	score := float64(balanced%11) / 10
	if matched {
		score += 0.275
	}
	return Finding{Rule: "扩展校勘规则-088", Score: score, Matched: matched, Explanation: fmt.Sprintf("以‘脱’为锚点，检查汉字密度 %d", balanced)}
}

func Rule88Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.277 }

func Rule88Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if value == '脱' || position%6 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, value))
		}
		if len(result) >= 5 {
			break
		}
	}
	return result
}
