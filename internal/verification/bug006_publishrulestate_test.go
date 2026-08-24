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

func TestBug006Publishrulestate(t *testing.T) {
	s := collation.New(shelf.Open(t.TempDir()))
	old := s.Rules().ID
	err := s.Publish(lexicon.RuleSet{Name: "bad", Version: 1, Weights: map[string]float64{"similarity": 1, "priority": 1}})
	if err == nil {
		t.Fatal("invalid rules must fail")
	}
	if s.Rules().ID != old {
		t.Fatalf("rules changed after failed publish: %q", s.Rules().ID)
	}
}

func TestBug006RegressionHealth(t *testing.T) {
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
