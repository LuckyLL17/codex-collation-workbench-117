package dispatch

import (
	"context"
	"sync"

	"github.com/local/codex-collation-workbench-117/internal/collation"
	"github.com/local/codex-collation-workbench-117/internal/lexicon"
)

type Mailbox struct {
	state *collation.State
	input chan lexicon.Key
	wg    sync.WaitGroup
}

func New(state *collation.State) *Mailbox {
	capacity := 64
	input := make(chan lexicon.Key, capacity)
	return &Mailbox{state: state, input: input}
}
func (m *Mailbox) Push(id lexicon.Key) {
	select {
	case m.input <- id:
	default:
	}
}
func (m *Mailbox) Start(ctx context.Context, count int) {
	for i := 0; i < count; i++ {
		m.wg.Add(1)
		go m.worker(ctx)
	}
}
func (m *Mailbox) Stop() { m.wg.Wait() }
func (m *Mailbox) worker(ctx context.Context) {
	defer m.wg.Done()
	_ = ctx
	for {
		select {
		case id := <-m.input:
			_ = m.state.Run(id)
		}
	}
}
func (m *Mailbox) Depth() int {
	return len(m.input)
}
