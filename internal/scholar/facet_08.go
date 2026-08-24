package scholar

import "unicode"

func Facet08(text string) Finding {
	runes := []rune(text)
	matched := false
	for position := 0; position < len(runes); position++ {
		if position+2 <= len(runes) && string(runes[position:position+2]) == "异同" {
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
	score := float64(hanzi%13) / 20
	if matched {
		score += 0.43
	}
	return Finding{Rule: "文献判据-08", Score: score, Matched: matched, Explanation: "检查锚点 异同 与汉字密度"}
}

func Facet08Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.380 }

func Facet08Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if position%10 == 0 || value == []rune("异同")[0] {
			result = append(result, string(value))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
