package verification

import (
	"context"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/local/codex-collation-workbench-117/internal/collation"
	"github.com/local/codex-collation-workbench-117/internal/dispatch"
	"github.com/local/codex-collation-workbench-117/internal/lexicon"
	"github.com/local/codex-collation-workbench-117/internal/presentation"
	"github.com/local/codex-collation-workbench-117/internal/runtime"
	"github.com/local/codex-collation-workbench-117/internal/scholar"
	"github.com/local/codex-collation-workbench-117/internal/shelf"
)

func TestBug010Priorityscoring(t *testing.T) {
	base := lexicon.Passage{Text: "相同"}
	high := lexicon.Score(base, lexicon.Passage{Text: "相同"}, lexicon.Witness{ID: "high", Priority: 9}, lexicon.DefaultRules())
	low := lexicon.Score(base, lexicon.Passage{Text: "相同"}, lexicon.Witness{ID: "low", Priority: 1}, lexicon.DefaultRules())
	if high.Score <= low.Score {
		t.Fatalf("high priority score %.3f <= low %.3f", high.Score, low.Score)
	}
}

func TestBug010RegressionHealth(t *testing.T) {
	if !dispatch.Available(1) {
		t.Fatal("health should remain available")
	}
}

var _ = context.Background
var _ = httptest.NewRecorder
var _ = os.ErrNotExist
var _ = filepath.Separator
var _ = time.Second
var _ = collation.ErrMissing
var _ = lexicon.Key("")
var _ = presentation.New
var _ = runtime.Load
var _ = scholar.HealthProfile
var _ = shelf.Open
