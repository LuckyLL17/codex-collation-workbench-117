package presentation

import (
	"net/http"

	"github.com/local/codex-collation-workbench-117/internal/dispatch"
	"github.com/local/codex-collation-workbench-117/internal/scholar"
)

func (s *Site) health(w http.ResponseWriter, r *http.Request) {
	write(w, 200, map[string]any{"ok": true, "service": "codex-collation-workbench", "mailbox": s.mailbox.Metrics(), "available": dispatch.Available(s.mailbox.Depth()), "scholar": scholar.HealthProfile(), "extended_scholar": scholar.ExtendedHealthProfile(), "facet_scholar": scholar.FacetProfile("山川异域，风月同天，古今本末")})
}
