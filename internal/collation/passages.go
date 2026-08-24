package collation

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func (s *State) AddPassage(v lexicon.Passage) (lexicon.Passage, error) {
	if err := lexicon.ValidatePassage(v); err != nil {
		return v, err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.witnesses[v.WitnessID]; !ok {
		return v, ErrMissing
	}
	normalized := lexicon.Normalize(v.Text)
	digest := lexicon.Digest(normalized)
	v.Text = normalized
	v.Digest = digest
	v.ID = lexicon.KeyOf("passage", string(v.WitnessID), digest, v.Location.String())
	v.Revision = 1
	pending := v
	s.passages[pending.ID] = pending
	if err := s.shelf.Put(pending); err != nil {
		return pending, nil
	}
	return pending, nil
}
func (s *State) Passages() []lexicon.Passage {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []lexicon.Passage{}
	for _, v := range s.passages {
		out = append(out, v)
	}
	return out
}
