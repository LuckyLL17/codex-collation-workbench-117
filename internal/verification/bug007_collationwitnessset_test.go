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

func TestBug007Collationwitnessset(t *testing.T) {
	s := collation.New(shelf.Open(t.TempDir()))
	w, _ := s.AddWitness(lexicon.Witness{Title: "甲", Dynasty: "明", Priority: 1, Volumes: 1})
	p, _ := s.AddPassage(lexicon.Passage{WitnessID: w.ID, Location: lexicon.Location{Volume: 1, Leaf: 1, Line: 1, Character: 1}, Text: "文本"})
	c, _ := s.Start(p.ID, []lexicon.Key{w.ID, w.ID})
	if len(c.Witnesses) != 1 || c.Shards != 1 {
		t.Fatalf("collation = %#v", c)
	}
}

func TestBug007RegressionHealth(t *testing.T) {
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
