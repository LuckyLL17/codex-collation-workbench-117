package verification

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/local/codex-collation-workbench-117/internal/collation"
	"github.com/local/codex-collation-workbench-117/internal/dispatch"
	"github.com/local/codex-collation-workbench-117/internal/lexicon"
	"github.com/local/codex-collation-workbench-117/internal/presentation"
	"github.com/local/codex-collation-workbench-117/internal/runtime"
	"github.com/local/codex-collation-workbench-117/internal/scholar"
	"github.com/local/codex-collation-workbench-117/internal/shelf"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestBug004Shelfmanifestcount(t *testing.T) {
	root := t.TempDir()
	s := shelf.Open(root)
	if err := s.Put(lexicon.Passage{Digest: lexicon.Digest("x"), Text: "x"}); err != nil {
		t.Fatal(err)
	}
	if got := s.ReadManifest().Blocks; got != 1 {
		t.Fatalf("manifest blocks = %d, want 1", got)
	}
}

func TestBug004RegressionHealth(t *testing.T) {
	if !dispatch.Available(0) {
		t.Fatal("health should remain available")
	}
}

var _ = context.Background
var _ = json.Valid
var _ = httptest.NewRecorder
var _ = os.ErrNotExist
var _ = filepath.Separator
var _ = time.Second
var _ = bytes.NewBuffer
var _ = collation.ErrMissing
var _ = lexicon.Key("")
var _ = presentation.New
var _ = runtime.Load
var _ = scholar.HealthProfile
var _ = shelf.Open
