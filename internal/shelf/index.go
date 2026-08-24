package shelf

import (
	"os"
	"path/filepath"
	"strings"
)

func (s *Shelf) Has(digest string) bool {
	_, err := os.Stat(filepath.Join(s.root, digest[:2], digest+".json"))
	return err == nil
}
func (s *Shelf) Prefix(prefix string) []string {
	out := []string{}
	entries, _ := os.ReadDir(s.root)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		files, _ := os.ReadDir(filepath.Join(s.root, entry.Name()))
		for _, file := range files {
			name := strings.TrimSuffix(file.Name(), ".json")
			if strings.HasPrefix(name, prefix) {
				out = append(out, name)
			}
		}
	}
	return out
}
