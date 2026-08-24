package main

import (
	"context"
	"errors"
	"github.com/local/codex-collation-workbench-117/internal/collation"
	"github.com/local/codex-collation-workbench-117/internal/dispatch"
	"github.com/local/codex-collation-workbench-117/internal/presentation"
	"github.com/local/codex-collation-workbench-117/internal/runtime"
	"github.com/local/codex-collation-workbench-117/internal/shelf"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	settings := runtime.Load()
	state := collation.New(shelf.Open(settings.Shelf))
	runtime.Seed(state)
	mailbox := dispatch.New(state)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	mailbox.Start(ctx, 3)
	server := &http.Server{Addr: settings.Address, Handler: presentation.New(state, mailbox).Handler()}
	go func() { <-ctx.Done(); mailbox.Stop(); _ = server.Shutdown(context.Background()) }()
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
