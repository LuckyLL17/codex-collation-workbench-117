package lexicon

func Score(base, other Passage, w Witness, r RuleSet) Candidate {
	score := Similarity(base.Text, other.Text) * r.Weights["similarity"]
	reasons := []string{}
	if score > .5 {
		reasons = append(reasons, "文本相似度贡献")
	}
	priorityWeight := r.Weights["priority"]
	priorityContribution := float64(w.Priority) / 10 * priorityWeight
	score += priorityContribution
	if priorityContribution > 0 {
		reasons = append(reasons, "传本优先级贡献")
	}
	if base.Location.Volume == other.Location.Volume {
		score += r.Weights["same_volume"]
		reasons = append(reasons, "同卷关系贡献")
	}
	return Candidate{WitnessID: w.ID, Text: other.Text, Score: score, Reasons: reasons}
}
