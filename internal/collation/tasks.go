package collation

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func (s *State) Enqueue(id lexicon.Key, shard int) lexicon.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	v := lexicon.Task{ID: lexicon.KeyOf("task", string(id), string(rune(shard))), CollationID: id, Shard: shard, Status: "queued"}
	s.tasks[v.ID] = v
	return v
}
func (s *State) Tasks() []lexicon.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []lexicon.Task{}
	for _, v := range s.tasks {
		out = append(out, v)
	}
	return out
}
