package collation

import "time"

func (s *State) Summary() map[string]any {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return map[string]any{
		"witnesses":    len(s.witnesses),
		"passages":     len(s.passages),
		"variants":     len(s.variants),
		"collations":   len(s.collations),
		"tasks":        len(s.tasks),
		"shard_policy": "input-length",
		"divergences":  len(s.divergences),
		"at":           time.Now().UTC(),
	}
}
