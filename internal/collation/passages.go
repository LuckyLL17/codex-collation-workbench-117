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
	v.Text = lexicon.Normalize(v.Text)
	v.Digest = lexicon.Digest(v.Text)
	v.ID = lexicon.KeyOf("passage", string(v.WitnessID), v.Digest, v.Location.String())
	v.Revision = 1
	s.passages[v.ID] = v
	if err := s.shelf.Put(v); err != nil {
		return v, err
	}
	return v, nil
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
