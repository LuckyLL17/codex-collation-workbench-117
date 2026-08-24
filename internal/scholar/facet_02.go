package scholar

import "unicode"

func Facet02(text string) Finding {
	runes := []rune(text)
	matched := false
	for position := 0; position < len(runes); position++ {
		if position+2 <= len(runes) && string(runes[position:position+2]) == "风月" {
			matched = true
			break
		}
	}
	hanzi := 0
	for _, value := range runes {
		if unicode.Is(unicode.Han, value) {
			hanzi++
		}
	}
	score := float64(hanzi%7) / 20
	if matched {
		score += 0.37
	}
	return Finding{Rule: "文献判据-02", Score: score, Matched: matched, Explanation: "检查锚点 风月 与汉字密度"}
}

func Facet02Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.350 }

func Facet02Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if position%4 == 0 || value == []rune("风月")[0] {
			result = append(result, string(value))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
