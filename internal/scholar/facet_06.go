package scholar

import "unicode"

func Facet06(text string) Finding {
	runes := []rune(text)
	matched := false
	for position := 0; position < len(runes); position++ {
		if position+2 <= len(runes) && string(runes[position:position+2]) == "卷帙" {
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
	score := float64(hanzi%11) / 20
	if matched {
		score += 0.41
	}
	return Finding{Rule: "文献判据-06", Score: score, Matched: matched, Explanation: "检查锚点 卷帙 与汉字密度"}
}

func Facet06Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.370 }

func Facet06Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if position%8 == 0 || value == []rune("卷帙")[0] {
			result = append(result, string(value))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
