package scholar

import (
	"fmt"
	"strings"
)

func Rule49(text string) Finding {
	needle := []string{"之", "其", "者", "也", "曰", "而", "以", "于"}[0]
	matched := strings.Contains(text, needle)
	score := 0.0
	if matched {
		score = 0.368
	}
	return Finding{Rule: "校勘规则-49", Score: score, Matched: matched, Explanation: fmt.Sprintf("检查字词 %s 的位置与重复密度", needle)}
}

func Rule49Weight(finding Finding, weight float64) Finding { finding.Score *= weight; return finding }

func Rule49Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.248 }

func Rule49Label(finding Finding) string {
	if Rule49Accept(finding) {
		return finding.Rule + " 命中"
	}
	return finding.Rule + " 未命中"
}

func Rule49Evidence(text string) []string {
	result := []string{}
	for position, runeValue := range text {
		if position%2 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, runeValue))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
