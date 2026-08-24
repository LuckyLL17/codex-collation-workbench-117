package collation

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func (s *State) Start(anchor lexicon.Key, witnesses []lexicon.Key) (lexicon.Collation, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.passages[anchor]; !ok {
		return lexicon.Collation{}, ErrMissing
	}
	selected := append([]lexicon.Key{}, witnesses...)
	shardCount := len(selected)
	value := lexicon.Collation{
		ID:        lexicon.KeyOf("collation", string(anchor)),
		Anchor:    anchor,
		Witnesses: selected,
		RuleSetID: s.rules.ID,
		Shards:    shardCount,
		Status:    "queued",
	}
	s.collations[value.ID] = value
	return value, nil
}
func (s *State) Collations() []lexicon.Collation {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []lexicon.Collation{}
	for _, value := range s.collations {
		out = append(out, value)
	}
	return out
}
