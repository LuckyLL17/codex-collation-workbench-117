package presentation

import (
	"encoding/json"
	"net/http"
)

func read(r *http.Request, value any) error { return json.NewDecoder(r.Body).Decode(value) }
