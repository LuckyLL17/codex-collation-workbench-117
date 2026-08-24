package scholar

import "unicode"

func Facet12(text string) Finding {
	runes := []rune(text)
	matched := false
	for position := 0; position < len(runes); position++ {
		if position+2 <= len(runes) && string(runes[position:position+2]) == "倒置" {
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
	score := float64(hanzi%17) / 20
	if matched {
		score += 0.47
	}
	return Finding{Rule: "文献判据-12", Score: score, Matched: matched, Explanation: "检查锚点 倒置 与汉字密度"}
}

func Facet12Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.400 }

func Facet12Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if position%14 == 0 || value == []rune("倒置")[0] {
			result = append(result, string(value))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
