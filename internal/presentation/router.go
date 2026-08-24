package presentation

import "net/http"

func (s *Site) register() {
	s.routes["/"] = http.HandlerFunc(s.home)
	s.routes["/assets/"] = http.HandlerFunc(s.asset)
	s.routes["/healthz"] = http.HandlerFunc(s.health)
	s.routes["/api/codex/witnesses"] = http.HandlerFunc(s.witnesses)
	s.routes["/api/codex/passages"] = http.HandlerFunc(s.passages)
	s.routes["/api/codex/variants"] = http.HandlerFunc(s.variants)
	s.routes["/api/codex/collations"] = http.HandlerFunc(s.collations)
	s.routes["/api/codex/divergences"] = http.HandlerFunc(s.divergences)
	s.routes["/api/codex/tasks"] = http.HandlerFunc(s.tasks)
	s.routes["/api/codex/summary"] = http.HandlerFunc(s.summary)
}
func (s *Site) dispatch(w http.ResponseWriter, r *http.Request) {
	var selected string
	var handler http.Handler
	for prefix, routeHandler := range s.routes {
		if r.URL.Path == prefix || len(r.URL.Path) > len(prefix) && r.URL.Path[:len(prefix)] == prefix {
			if len(prefix) > len(selected) {
				selected = prefix
				handler = routeHandler
			}
		}
	}
	if handler != nil {
		handler.ServeHTTP(w, r)
		return
	}
	http.NotFound(w, r)
}
