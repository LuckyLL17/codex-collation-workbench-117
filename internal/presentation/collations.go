package presentation

import (
	"github.com/local/codex-collation-workbench-117/internal/lexicon"
	"net/http"
)

func (s *Site) collations(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		write(w, 200, s.state.Collations())
		return
	}
	if r.Method != http.MethodPost {
		badMethod(w)
		return
	}
	var input struct {
		Anchor    lexicon.Key   `json:"anchor"`
		Witnesses []lexicon.Key `json:"witnesses"`
	}
	if err := read(r, &input); err != nil {
		write(w, 400, map[string]any{"error": err.Error()})
		return
	}
	value, err := s.state.Start(input.Anchor, input.Witnesses)
	if err != nil {
		write(w, 422, map[string]any{"error": err.Error()})
		return
	}
	for shard := 0; shard < value.Shards; shard++ {
		task := s.state.Enqueue(value.ID, shard)
		s.mailbox.Push(task.ID)
	}
	write(w, 202, value)
}
