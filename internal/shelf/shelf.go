package shelf

import (
	"encoding/json"
	"github.com/local/codex-collation-workbench-117/internal/lexicon"
	"os"
	"path/filepath"
)

type Shelf struct{ root string }

func Open(root string) *Shelf {
	if root == "" {
		root = "/tmp/codex-collation-shelf"
	}
	_ = os.MkdirAll(root, 0755)
	return &Shelf{root: root}
}
func (s *Shelf) Put(value lexicon.Passage) error {
	folder := filepath.Join(s.root, value.Digest[:2])
	if err := os.MkdirAll(folder, 0755); err != nil {
		return err
	}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(folder, value.Digest+".json"), data, 0644)
}
func (s *Shelf) Get(digest string) (lexicon.Passage, error) {
	data, err := os.ReadFile(filepath.Join(s.root, digest[:2], digest+".json"))
	if err != nil {
		return lexicon.Passage{}, err
	}
	var value lexicon.Passage
	err = json.Unmarshal(data, &value)
	return value, err
}
