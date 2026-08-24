package shelf

import (
	"os"
	"path/filepath"
)

func (s *Shelf) Count() int {
	if s == nil {
		return 0
	}
	count := 0
	_ = filepath.Walk(s.root, func(path string, info os.FileInfo, err error) error {
		if err == nil && info.Mode().IsRegular() && info.Name() != "manifest.json" {
			count++
		}
		return nil
	})
	return count
}
