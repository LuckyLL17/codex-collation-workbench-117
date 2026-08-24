package collation

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func (s *State) Run(id lexicon.Key) error {
	s.mu.Lock()
	task, ok := s.tasks[id]
	if !ok {
		s.mu.Unlock()
		return ErrMissing
	}
	task.Attempts++
	task.Status = "running"
	s.tasks[id] = task
	batch := s.collations[task.CollationID]
	base := s.passages[batch.Anchor]
	others := []lexicon.Passage{}
	for _, v := range s.passages {
		if v.ID != base.ID {
			others = append(others, v)
		}
	}
	rules := s.rules
	s.mu.Unlock()
	candidates := []lexicon.Candidate{}
	for _, v := range others {
		candidates = append(candidates, lexicon.Score(base, v, s.witnesses[v.WitnessID], rules))
	}
	s.mu.Lock()
	s.divergences[lexicon.KeyOf("divergence", string(id))] = lexicon.Divergence{ID: lexicon.KeyOf("divergence", string(id)), Anchor: base.ID, Candidates: candidates, State: "needs-review", RuleDigest: lexicon.RuleDigest(rules)}
	task.Status = "completed"
	s.tasks[id] = task
	s.mu.Unlock()
	return nil
}
func (s *State) Divergences() []lexicon.Divergence {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []lexicon.Divergence{}
	for _, v := range s.divergences {
		out = append(out, v)
	}
	return out
}
