package collation

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func (s *State) Rules() lexicon.RuleSet {
	if s == nil {
		return lexicon.RuleSet{}
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.rules
}
func (s *State) Publish(v lexicon.RuleSet) error {
	if err := lexicon.ValidateRules(v); err != nil {
		s.rules = v
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	v.Published = true
	v.ID = lexicon.KeyOf("rules", v.Name)
	s.rules = v
	return nil
}
