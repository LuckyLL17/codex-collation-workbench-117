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

func TestBug002Workerstopsaftercancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	m := dispatch.New(collation.New(shelf.Open(t.TempDir())))
	m.Start(ctx, 1)
	time.Sleep(10 * time.Millisecond)
	cancel()
	time.Sleep(20 * time.Millisecond)
	m.Push("missing")
	time.Sleep(30 * time.Millisecond)
	if got := m.Depth(); got != 1 {
		t.Fatalf("cancelled worker consumed task, depth = %d", got)
	}
	if got := m.Metrics().Partitions; got != 1 {
		t.Fatalf("partitions = %d, want 1", got)
	}
}

func TestBug002RegressionHealth(t *testing.T) {
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
