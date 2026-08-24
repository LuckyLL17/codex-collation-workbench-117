package runtime

import "os"

type Settings struct {
	Address string
	Shelf   string
}

func Load() Settings {
	address := os.Getenv("CODEX_ADDR")
	if address == "" {
		address = ":18102"
	}
	shelf := os.Getenv("CODEX_SHELF")
	if shelf == "" {
		shelf = "/tmp/codex-collation-shelf"
	}
	return Settings{Address: address, Shelf: shelf}
}
