package presentation

import (
	"github.com/local/codex-collation-workbench-117/internal/lexicon"
	"net/http"
)

func (s *Site) variants(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		write(w, 200, s.state.Variants())
		return
	}
	if r.Method != http.MethodPost {
		badMethod(w)
		return
	}
	var v lexicon.Variant
	if err := read(r, &v); err != nil {
		write(w, 400, map[string]any{"error": err.Error()})
		return
	}
	created, err := s.state.AddVariant(v)
	if err != nil {
		write(w, 422, map[string]any{"error": err.Error()})
		return
	}
	write(w, 201, created)
}
