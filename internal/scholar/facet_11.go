package scholar

import "unicode"

func Facet11(text string) Finding {
	runes := []rune(text)
	matched := false
	for position := 0; position < len(runes); position++ {
		if position+2 <= len(runes) && string(runes[position:position+2]) == "脱衍" {
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
	score := float64(hanzi%16) / 20
	if matched {
		score += 0.46
	}
	return Finding{Rule: "文献判据-11", Score: score, Matched: matched, Explanation: "检查锚点 脱衍 与汉字密度"}
}

func Facet11Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.395 }

func Facet11Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if position%13 == 0 || value == []rune("脱衍")[0] {
			result = append(result, string(value))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
