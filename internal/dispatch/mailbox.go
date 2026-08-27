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

	// closed guards the lifecycle of the input channel: once true, Push is a
	// no-op so no new tasks enter the pipeline after shutdown begins. The mutex
	// also serializes Push against the channel close in Stop, which is what
	// makes sending on a possibly-closing channel panic-free.
	mu     sync.Mutex
	closed bool

	// partitions records the number of background consumers started for the
	// current run, so the status surface reports an accurate worker count.
	partitions int
}

func New(state *collation.State) *Mailbox {
	capacity := 64
	input := make(chan lexicon.Key, capacity)
	return &Mailbox{state: state, input: input}
}
func (m *Mailbox) Push(id lexicon.Key) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.closed {
		// Shutdown is in flight: drop the task rather than enqueue work that
		// would be processed after the stop signal has been issued.
		return
	}
	select {
	case m.input <- id:
	default:
	}
}
func (m *Mailbox) Start(ctx context.Context, count int) {
	m.mu.Lock()
	m.partitions = count
	m.mu.Unlock()
	for i := 0; i < count; i++ {
		m.wg.Add(1)
		go m.worker(ctx)
	}
}
func (m *Mailbox) Stop() {
	m.mu.Lock()
	if !m.closed {
		m.closed = true
		close(m.input)
	}
	m.mu.Unlock()
	m.wg.Wait()
}
func (m *Mailbox) worker(ctx context.Context) {
	defer m.wg.Done()
	for {
		select {
		case <-ctx.Done():
			// Stop signal received: stop pulling new tasks immediately.
			return
		case id, ok := <-m.input:
			if !ok {
				// Channel was closed by Stop; drain complete.
				return
			}
			if ctx.Err() != nil {
				// The task was dequeued concurrently with cancellation. Do not
				// process it: honor the stop signal over this late arrival.
				continue
			}
			_ = m.state.Run(id)
		}
	}
}
func (m *Mailbox) Depth() int {
	// Safe to call during and after shutdown; len on a closed channel reports
	// the remaining buffered items, so the status surface stays available.
	return len(m.input)
}
func (m *Mailbox) closedState() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.closed
}
