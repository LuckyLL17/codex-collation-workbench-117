package scholar

import "unicode"

func Facet04(text string) Finding {
	runes := []rune(text)
	matched := false
	for position := 0; position < len(runes); position++ {
		if position+2 <= len(runes) && string(runes[position:position+2]) == "本末" {
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
	score := float64(hanzi%9) / 20
	if matched {
		score += 0.39
	}
	return Finding{Rule: "文献判据-04", Score: score, Matched: matched, Explanation: "检查锚点 本末 与汉字密度"}
}

func Facet04Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.360 }

func Facet04Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if position%6 == 0 || value == []rune("本末")[0] {
			result = append(result, string(value))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
