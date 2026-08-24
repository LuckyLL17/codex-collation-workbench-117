package collation

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func ReviewState(value lexicon.Divergence) string {
	if len(value.Candidates) == 0 {
		return "empty"
	}
	if lexicon.Best(value.Candidates).Score > .8 {
		return "strong-candidate"
	}
	return value.State
}
