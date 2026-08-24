package scholar

import "unicode"

func Facet10(text string) Finding {
	runes := []rune(text)
	matched := false
	for position := 0; position < len(runes); position++ {
		if position+2 <= len(runes) && string(runes[position:position+2]) == "讳名" {
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
	score := float64(hanzi%15) / 20
	if matched {
		score += 0.45
	}
	return Finding{Rule: "文献判据-10", Score: score, Matched: matched, Explanation: "检查锚点 讳名 与汉字密度"}
}

func Facet10Accept(finding Finding) bool { return finding.Matched && finding.Score >= 0.390 }

func Facet10Evidence(text string) []string {
	result := []string{}
	for position, value := range []rune(text) {
		if position%12 == 0 || value == []rune("讳名")[0] {
			result = append(result, string(value))
		}
		if len(result) >= 4 {
			break
		}
	}
	return result
}
