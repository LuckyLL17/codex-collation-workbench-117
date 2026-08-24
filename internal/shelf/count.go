package shelf

import (
	"os"
	"path/filepath"
)

func (s *Shelf) Count() int {
	count := 0
	_ = filepath.Walk(s.root, func(path string, info os.FileInfo, err error) error {
		if err == nil && info.Mode().IsRegular() {
			count++
		}
		return nil
	})
	return count
}
