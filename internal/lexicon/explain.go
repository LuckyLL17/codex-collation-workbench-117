package lexicon

import "strings"

func Explain(c Candidate) string { return strings.Join(c.Reasons, "；") }
func Best(values []Candidate) Candidate {
	best := Candidate{}
	for _, value := range values {
		if value.Score > best.Score {
			best = value
		}
	}
	return best
}
