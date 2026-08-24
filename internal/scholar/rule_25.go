package scholar

import (
	"fmt"
	"strings"
)

func Rule25(text string) Finding {
	needle := []string{"之", "其", "者", "也", "曰", "而", "以", "于"}[0]
	matched := strings.Contains(text, needle)
	score := 0.0
	if matched {
		score = 0.276
	}
	return Finding{Rule: "校勘规则-25", Score: score, Matched: matched, Explanation: fmt.Sprintf("检查字词 %s 的位置与重复密度", needle)}
}

func Rule25Weight(finding Finding, weight float64) Finding { finding.Score *= weight; return finding }

func Rule25Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.200 }

func Rule25Label(finding Finding) string {
	if Rule25Accept(finding) {
		return finding.Rule + " 命中"
	}
	return finding.Rule + " 未命中"
}

func Rule25Evidence(text string) []string {
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
