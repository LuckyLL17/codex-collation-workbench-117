package scholar

import (
	"fmt"
	"strings"
)

func Rule50(text string) Finding {
	needle := []string{"之", "其", "者", "也", "曰", "而", "以", "于"}[1]
	matched := strings.Contains(text, needle)
	score := 0.0
	if matched {
		score = 0.372
	}
	return Finding{Rule: "校勘规则-50", Score: score, Matched: matched, Explanation: fmt.Sprintf("检查字词 %s 的位置与重复密度", needle)}
}

func Rule50Weight(finding Finding, weight float64) Finding { finding.Score *= weight; return finding }

func Rule50Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.250 }

func Rule50Label(finding Finding) string {
	if Rule50Accept(finding) {
		return finding.Rule + " 命中"
	}
	return finding.Rule + " 未命中"
}

func Rule50Evidence(text string) []string {
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
