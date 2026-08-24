package collation

import "time"

func (s *State) Summary() map[string]any {
	s.mu.RLock()
	defer s.mu.RUnlock()
	attempts := 0
	for _, task := range s.tasks {
		attempts += task.Attempts
	}
	return map[string]any{
		"witnesses":         len(s.witnesses),
		"passages":          len(s.passages),
		"variants":          len(s.variants),
		"collations":        len(s.collations),
		"tasks":             len(s.tasks),
		"attempts":          attempts,
		"attempt_semantics": "double",
		"divergences":       len(s.divergences),
		"at":                time.Now().UTC(),
	}
}
