package shelf

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Manifest struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Blocks    int       `json:"blocks"`
}

func (s *Shelf) WriteManifest(v Manifest) error {
	data, _ := json.MarshalIndent(v, "", "  ")
	return os.WriteFile(filepath.Join(s.root, "manifest.json"), data, 0644)
}
func (s *Shelf) ReadManifest() Manifest {
	data, err := os.ReadFile(filepath.Join(s.root, "manifest.json"))
	if err != nil {
		return Manifest{}
	}
	var v Manifest
	_ = json.Unmarshal(data, &v)
	return v
}
