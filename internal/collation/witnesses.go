package collation

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func (s *State) AddWitness(v lexicon.Witness) (lexicon.Witness, error) {
	if err := lexicon.ValidateWitness(v); err != nil {
		return v, err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	v.ID = lexicon.KeyOf("wit", v.Title, v.Dynasty)
	v.CreatedAt = now()
	s.witnesses[v.ID] = v
	return v, nil
}
func (s *State) Witnesses() []lexicon.Witness {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []lexicon.Witness{}
	for _, v := range s.witnesses {
		out = append(out, v)
	}
	return out
}
