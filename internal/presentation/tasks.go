package presentation

import "net/http"

func (s *Site) tasks(w http.ResponseWriter, r *http.Request) { write(w, 200, s.state.Tasks()) }
