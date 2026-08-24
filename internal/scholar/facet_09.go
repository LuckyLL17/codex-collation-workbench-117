package scholar

import "unicode"

func Facet09(text string) Finding {
	runes := []rune(text)
	matched := false
	for position := 0; position < len(runes); position++ {
		if position+2 <= len(runes) && string(runes[position:position+2]) == "音训" {
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
	score := float64(hanzi%14) / 20
	if matched {
		score += 0.44
	}
	return Finding{Rule: "文献判据-09", Score: score, Matched: matched, Explanation: "检查锚点 音训 与汉字密度"}
}

func Facet09Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.385 }

func Facet09Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if position%11 == 0 || value == []rune("音训")[0] {
			result = append(result, string(value))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
