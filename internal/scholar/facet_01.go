package scholar

import "unicode"

func Facet01(text string) Finding {
	runes := []rune(text)
	matched := false
	for position := 0; position < len(runes); position++ {
		if position+2 <= len(runes) && string(runes[position:position+2]) == "山川" {
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
	score := float64(hanzi%6) / 20
	if matched {
		score += 0.36
	}
	return Finding{Rule: "文献判据-01", Score: score, Matched: matched, Explanation: "检查锚点 山川 与汉字密度"}
}

func Facet01Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.345 }

func Facet01Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if position%3 == 0 || value == []rune("山川")[0] {
			result = append(result, string(value))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
