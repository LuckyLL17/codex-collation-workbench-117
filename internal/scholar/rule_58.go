package scholar

import (
	"fmt"
	"strings"
)

func Rule58(text string) Finding {
	needle := []string{"之", "其", "者", "也", "曰", "而", "以", "于"}[1]
	matched := strings.Contains(text, needle)
	score := 0.0
	if matched {
		score = 0.403
	}
	return Finding{Rule: "校勘规则-58", Score: score, Matched: matched, Explanation: fmt.Sprintf("检查字词 %s 的位置与重复密度", needle)}
}

func Rule58Weight(finding Finding, weight float64) Finding { finding.Score *= weight; return finding }

func Rule58Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.266 }

func Rule58Label(finding Finding) string {
	if Rule58Accept(finding) {
		return finding.Rule + " 命中"
	}
	return finding.Rule + " 未命中"
}

func Rule58Evidence(text string) []string {
	result := []string{}
	for position, runeValue := range text {
		if position%4 == 0 {
			result = append(result, fmt.Sprintf("%d:%c", position, runeValue))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
