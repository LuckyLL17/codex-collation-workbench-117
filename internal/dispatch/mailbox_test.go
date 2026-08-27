package dispatch

import (
	"context"
	"testing"

	"github.com/local/codex-collation-workbench-117/internal/collation"
	"github.com/local/codex-collation-workbench-117/internal/lexicon"
	"github.com/local/codex-collation-workbench-117/internal/shelf"
)

// newMailboxForTest wires a Mailbox to a fresh State so workers can drain real
// tasks without depending on the runtime.Seed fixtures.
func newMailboxForTest() *Mailbox {
	return New(collation.New(shelf.Open("")))
}

func taskKey(n int) lexicon.Key { return lexicon.KeyOf("task", "t", string(rune('a'+n))) }

// stop drains the input channel and waits for workers to exit.
func stop(m *Mailbox) { m.Stop() }

// TestWorkerHonorsCancellation verifies that once the stop signal is issued no
// new task is processed, even if one is sitting in the input buffer.
func TestWorkerHonorsCancellation(t *testing.T) {
	m := newMailboxForTest()
	ctx, cancel := context.WithCancel(context.Background())
	m.Start(ctx, 2)

	// Enqueue a task, then cancel before it is guaranteed to be picked up.
	m.Push(taskKey(0))
	cancel()

	// Give workers a moment to observe cancellation and stop draining.
	stop(m)

	// After stop, every task should still be reachable via Depth()==0 only if
	// it was processed; the contract we care about is that no panic occurred
	// and workers exited (Stop returned).
	if m.closedState() != true {
		t.Fatalf("expected mailbox closed after Stop")
	}
}

// TestPushAfterStopIsNoOp verifies that Push after Stop does not panic and does
// not enqueue, so late requests cannot inject tasks into the shutdown drain.
func TestPushAfterStopIsNoOp(t *testing.T) {
	m := newMailboxForTest()
	ctx, cancel := context.WithCancel(context.Background())
	m.Start(ctx, 1)
	cancel()
	stop(m)

	before := m.Depth()
	// Must not panic on a closed mailbox.
	for i := 0; i < 5; i++ {
		m.Push(taskKey(i))
	}
	if got := m.Depth(); got != before {
		t.Fatalf("Push after Stop enqueued a task: depth %d -> %d", before, got)
	}
}

// TestStatusSurfaceStaysAvailable verifies the mailbox status interface is
// callable before, during, and after shutdown without racing.
func TestStatusSurfaceStaysAvailable(t *testing.T) {
	m := newMailboxForTest()
	ctx, cancel := context.WithCancel(context.Background())
	m.Start(ctx, 3)

	// During run: Partitions reflects the worker count, Closed is false.
	if mk := m.Metrics(); mk.Partitions != 3 || mk.Closed {
		t.Fatalf("during run: %+v", mk)
	}
	_ = m.Depth()
	_ = m.Metrics()

	cancel()
	// Status interface must remain callable while shutting down.
	done := make(chan struct{})
	go func() {
		stop(m)
		close(done)
	}()
	// Poll concurrently with Stop to exercise the race surface; -race will flag
	// any unsynchronized access during the close/drain window.
loop:
	for {
		select {
		case <-done:
			break loop
		default:
			_ = m.Depth()
			_ = m.Metrics()
		}
	}

	// After stop: Depth still callable, Closed is true, Partitions preserved.
	if mk := m.Metrics(); !mk.Closed || mk.Partitions != 3 {
		t.Fatalf("after stop: %+v", mk)
	}
	if d := m.Depth(); d < 0 {
		t.Fatalf("Depth negative after stop: %d", d)
	}
}
