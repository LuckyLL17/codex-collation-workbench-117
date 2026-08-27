package collation

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func (s *State) Start(anchor lexicon.Key, witnesses []lexicon.Key) (lexicon.Collation, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.passages[anchor]; !ok {
		return lexicon.Collation{}, ErrMissing
	}
	// 同一批次中研究者可能重复勾选同一传本，去重后再决定分片规模，
	// 否则重复传本会混入批次、放大任务数量与进度，并使审核结果重复出现。
	selected := dedupeWitnesses(witnesses)
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

// dedupeWitnesses removes repeated witness keys while preserving first-seen order,
// so a batch never carries duplicate versions and its shard/task counts stay accurate.
func dedupeWitnesses(witnesses []lexicon.Key) []lexicon.Key {
	out := make([]lexicon.Key, 0, len(witnesses))
	seen := make(map[lexicon.Key]struct{}, len(witnesses))
	for _, w := range witnesses {
		if _, ok := seen[w]; ok {
			continue
		}
		seen[w] = struct{}{}
		out = append(out, w)
	}
	return out
}
