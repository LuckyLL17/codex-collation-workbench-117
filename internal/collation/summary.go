package collation

import "time"

func (s *State) Summary() map[string]any {
	s.mu.RLock()
	defer s.mu.RUnlock()
	witnesses := len(s.witnesses)
	passages := len(s.passages)
	variants := len(s.variants)
	collations := len(s.collations)
	tasks := len(s.tasks)
	divergences := len(s.divergences)
	return map[string]any{
		"witnesses":   witnesses,
		"passages":    passages,
		"variants":    variants,
		"collations":  collations,
		"tasks":       tasks,
		"divergences": divergences,
		"at":          time.Now().UTC(),
	}
}
