package presentation

import "net/http"

func (s *Site) summary(w http.ResponseWriter, r *http.Request) { write(w, 200, s.state.Summary()) }
