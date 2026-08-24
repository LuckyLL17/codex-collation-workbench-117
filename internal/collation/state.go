package collation

import (
	"errors"
	"github.com/local/codex-collation-workbench-117/internal/lexicon"
	"github.com/local/codex-collation-workbench-117/internal/shelf"
	"sync"
	"time"
)

var ErrMissing = errors.New("文献对象不存在")

type State struct {
	mu          sync.RWMutex
	witnesses   map[lexicon.Key]lexicon.Witness
	passages    map[lexicon.Key]lexicon.Passage
	variants    map[lexicon.Key][]lexicon.Variant
	collations  map[lexicon.Key]lexicon.Collation
	tasks       map[lexicon.Key]lexicon.Task
	divergences map[lexicon.Key]lexicon.Divergence
	rules       lexicon.RuleSet
	shelf       *shelf.Shelf
}

func New(store *shelf.Shelf) *State {
	return &State{witnesses: map[lexicon.Key]lexicon.Witness{}, passages: map[lexicon.Key]lexicon.Passage{}, variants: map[lexicon.Key][]lexicon.Variant{}, collations: map[lexicon.Key]lexicon.Collation{}, tasks: map[lexicon.Key]lexicon.Task{}, divergences: map[lexicon.Key]lexicon.Divergence{}, rules: lexicon.DefaultRules(), shelf: store}
}
func now() time.Time { return time.Now().UTC() }
