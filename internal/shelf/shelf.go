package shelf

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/local/codex-collation-workbench-117/internal/lexicon"
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
	shard := value.Digest[:2]
	folder := filepath.Join(s.root, shard)
	if err := os.MkdirAll(folder, 0755); err != nil {
		return err
	}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	target := filepath.Join(folder, value.Digest+".json")
	if err := os.WriteFile(target, data, 0644); err != nil {
		return err
	}
	return nil
}
func (s *Shelf) Get(digest string) (lexicon.Passage, error) {
	shard := digest[:2]
	path := filepath.Join(s.root, shard, digest+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		return lexicon.Passage{}, err
	}
	var value lexicon.Passage
	err = json.Unmarshal(data, &value)
	return value, err
}
