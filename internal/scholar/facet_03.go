package scholar

import "unicode"

func Facet03(text string) Finding {
	runes := []rune(text)
	matched := false
	for position := 0; position < len(runes); position++ {
		if position+2 <= len(runes) && string(runes[position:position+2]) == "古今" {
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
	score := float64(hanzi%8) / 20
	if matched {
		score += 0.38
	}
	return Finding{Rule: "文献判据-03", Score: score, Matched: matched, Explanation: "检查锚点 古今 与汉字密度"}
}

func Facet03Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.355 }

func Facet03Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if position%5 == 0 || value == []rune("古今")[0] {
			result = append(result, string(value))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
