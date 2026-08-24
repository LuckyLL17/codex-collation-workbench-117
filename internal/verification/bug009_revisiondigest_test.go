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

func TestBug009Revisiondigest(t *testing.T) {
	v := lexicon.Passage{Text: "  山川  "}
	v = lexicon.NextRevision(v)
	if v.Digest != lexicon.Digest(v.Text) {
		t.Fatalf("digest %q does not match text", v.Digest)
	}
}

func TestBug009RegressionHealth(t *testing.T) {
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
