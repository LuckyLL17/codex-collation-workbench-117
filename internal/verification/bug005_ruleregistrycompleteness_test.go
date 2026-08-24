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

func TestBug005Ruleregistrycompleteness(t *testing.T) {
	profile := scholar.HealthProfile()
	if profile["rule_count"] != 65 {
		t.Fatalf("rule count = %#v, want 65", profile["rule_count"])
	}
}

func TestBug005RegressionHealth(t *testing.T) {
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
