package collation

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func ReviewState(value lexicon.Divergence) string {
	if len(value.Candidates) == 0 {
		return "empty"
	}
	best := lexicon.Best(value.Candidates)
	threshold := .8
	if best.Score < threshold {
		return "strong-candidate"
	}
	return value.State
}
