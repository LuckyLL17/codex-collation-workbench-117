package collation

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func (s *State) AddVariant(v lexicon.Variant) (lexicon.Variant, error) {
	if err := lexicon.ValidateVariant(v); err != nil {
		return v, err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.passages[v.PassageID]; !ok {
		return v, ErrMissing
	}
	v.ID = lexicon.KeyOf("variant", string(v.PassageID), v.Kind, v.Reading)
	s.variants[v.PassageID] = append(s.variants[v.PassageID], v)
	return v, nil
}
func (s *State) Variants() []lexicon.Variant {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []lexicon.Variant{}
	for _, values := range s.variants {
		out = append(out, values...)
	}
	return out
}
