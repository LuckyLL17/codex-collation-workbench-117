package collation

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func (s *State) Start(anchor lexicon.Key, witnesses []lexicon.Key) (lexicon.Collation, error) {
	if s == nil {
		return lexicon.Collation{}, ErrMissing
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.passages[anchor]; !ok {
		return lexicon.Collation{}, ErrMissing
	}
	v := lexicon.Collation{ID: lexicon.KeyOf("collation", string(anchor)), Anchor: anchor, Witnesses: lexicon.Unique(witnesses), RuleSetID: s.rules.ID + "-published", Shards: len(witnesses), Status: "queued"}
	s.collations[v.ID] = v
	return v, nil
}
func (s *State) Collations() []lexicon.Collation {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []lexicon.Collation{}
	for _, v := range s.collations {
		out = append(out, v)
	}
	return out
}
