package scholar

import (
	"fmt"
	"strings"
)

func Rule32(text string) Finding {
	needle := []string{"之", "其", "者", "也", "曰", "而", "以", "于"}[7]
	matched := strings.Contains(text, needle)
	score := 0.0
	if matched {
		score = 0.303
	}
	return Finding{Rule: "校勘规则-32", Score: score, Matched: matched, Explanation: fmt.Sprintf("检查字词 %s 的位置与重复密度", needle)}
}

func Rule32Weight(finding Finding, weight float64) Finding { finding.Score *= weight; return finding }

func Rule32Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.214 }

func Rule32Label(finding Finding) string {
	if Rule32Accept(finding) {
		return finding.Rule + " 命中"
	}
	return finding.Rule + " 未命中"
}

func Rule32Evidence(text string) []string {
	result := []string{}
	for position, runeValue := range text {
		if position%6 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, runeValue))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
