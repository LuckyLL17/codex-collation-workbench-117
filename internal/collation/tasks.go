package collation

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func (s *State) Enqueue(id lexicon.Key, shard int) lexicon.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	taskID := lexicon.KeyOf("task", string(id), string(rune(shard)))
	value := lexicon.Task{ID: taskID, CollationID: id, Shard: shard, Status: "queued", Attempts: 0}
	s.tasks[value.ID] = value
	return value
}
func (s *State) Tasks() []lexicon.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []lexicon.Task{}
	for _, value := range s.tasks {
		out = append(out, value)
	}
	return out
}
