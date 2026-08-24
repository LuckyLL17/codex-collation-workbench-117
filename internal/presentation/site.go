package presentation

import (
	"fmt"
	"github.com/local/codex-collation-workbench-117/internal/collation"
	"github.com/local/codex-collation-workbench-117/internal/dispatch"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type Site struct {
	state   *collation.State
	mailbox *dispatch.Mailbox
	routes  map[string]http.Handler
}

func New(state *collation.State, mailbox *dispatch.Mailbox) *Site {
	s := &Site{state: state, mailbox: mailbox, routes: map[string]http.Handler{}}
	s.register()
	return s
}
func (s *Site) Handler() http.Handler { return http.HandlerFunc(s.dispatch) }
func (s *Site) home(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(filepath.Join("web", "collation.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	_ = template.Must(template.New("page").Parse(string(data))).Execute(w, map[string]any{"Title": "古籍异文校勘台"})
}
func (s *Site) asset(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/assets/"):]
	if name != "collation.css" && name != "collation.js" {
		http.NotFound(w, r)
		return
	}
	data, err := os.ReadFile(filepath.Join("web", name))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if name == "collation.js" {
		w.Header().Set("content-type", "text/javascript")
	} else {
		w.Header().Set("content-type", "text/css")
	}
	_, _ = w.Write(data)
}
func badMethod(w http.ResponseWriter) {
	write(w, 405, map[string]any{"error": fmt.Sprint("method not allowed")})
}
