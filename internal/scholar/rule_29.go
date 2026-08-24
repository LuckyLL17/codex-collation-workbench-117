package scholar

import (
	"fmt"
	"strings"
)

func Rule29(text string) Finding {
	needle := []string{"之", "其", "者", "也", "曰", "而", "以", "于"}[4]
	matched := strings.Contains(text, needle)
	score := 0.0
	if matched {
		score = 0.292
	}
	return Finding{Rule: "校勘规则-29", Score: score, Matched: matched, Explanation: fmt.Sprintf("检查字词 %s 的位置与重复密度", needle)}
}

func Rule29Weight(finding Finding, weight float64) Finding { finding.Score *= weight; return finding }

func Rule29Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.208 }

func Rule29Label(finding Finding) string {
	if Rule29Accept(finding) {
		return finding.Rule + " 命中"
	}
	return finding.Rule + " 未命中"
}

func Rule29Evidence(text string) []string {
	result := []string{}
	for position, runeValue := range text {
		if position%3 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, runeValue))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
