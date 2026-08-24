package runtime

import (
	"github.com/local/codex-collation-workbench-117/internal/collation"
	"github.com/local/codex-collation-workbench-117/internal/lexicon"
)

func Seed(state *collation.State) {
	a, _ := state.AddWitness(lexicon.Witness{Title: "甲本", Dynasty: "明", Priority: 5, Volumes: 2})
	b, _ := state.AddWitness(lexicon.Witness{Title: "乙本", Dynasty: "清", Priority: 3, Volumes: 2})
	p, _ := state.AddPassage(lexicon.Passage{WitnessID: a.ID, Location: lexicon.Location{Volume: 1, Leaf: 2, Line: 3, Character: 1}, Text: "山川异域 风月同天"})
	_, _ = state.AddPassage(lexicon.Passage{WitnessID: b.ID, Location: lexicon.Location{Volume: 1, Leaf: 2, Line: 3, Character: 1}, Text: "山河异域 风月同天"})
	_, _ = state.AddVariant(lexicon.Variant{PassageID: p.ID, Kind: "形近", Reading: "河", Note: "山川与山河"})
}
