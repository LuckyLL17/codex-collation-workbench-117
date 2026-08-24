package collation

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func (s *State) Passage(id lexicon.Key) (lexicon.Passage, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.passages[id]
	return v, ok
}
func (s *State) Witness(id lexicon.Key) (lexicon.Witness, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.witnesses[id]
	return v, ok
}
