package presentation

import (
	"github.com/local/codex-collation-workbench-117/internal/lexicon"
	"net/http"
)

func (s *Site) witnesses(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		write(w, 200, s.state.Witnesses())
		return
	}
	if r.Method != http.MethodPost {
		badMethod(w)
		return
	}
	var v lexicon.Witness
	if err := read(r, &v); err != nil {
		write(w, 400, map[string]any{"error": err.Error()})
		return
	}
	created, err := s.state.AddWitness(v)
	if err != nil {
		write(w, 422, map[string]any{"error": err.Error()})
		return
	}
	write(w, 201, created)
}
