package scholar

import "unicode"

func Facet07(text string) Finding {
	runes := []rune(text)
	matched := false
	for position := 0; position < len(runes); position++ {
		if position+2 <= len(runes) && string(runes[position:position+2]) == "校定" {
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
	score := float64(hanzi%12) / 20
	if matched {
		score += 0.42
	}
	return Finding{Rule: "文献判据-07", Score: score, Matched: matched, Explanation: "检查锚点 校定 与汉字密度"}
}

func Facet07Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.375 }

func Facet07Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if position%9 == 0 || value == []rune("校定")[0] {
			result = append(result, string(value))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
