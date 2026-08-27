package collation

import (
	"sync"
	"testing"

	"github.com/local/codex-collation-workbench-117/internal/lexicon"
	"github.com/local/codex-collation-workbench-117/internal/shelf"
)

// TestRunConcurrentAttemptConsistency 模拟两个后台执行操作交叠触发同一条校勘任务：
// 实际执行体应只完成一次，进度统计的尝试次数也应只增加一次，避免监控误判为发生重试。
func TestRunConcurrentAttemptConsistency(t *testing.T) {
	store := shelf.Open(t.TempDir())
	state := New(store)

	a, err := state.AddWitness(lexicon.Witness{Title: "甲本", Dynasty: "明", Priority: 5, Volumes: 2})
	if err != nil {
		t.Fatalf("seed witness a: %v", err)
	}
	b, err := state.AddWitness(lexicon.Witness{Title: "乙本", Dynasty: "清", Priority: 3, Volumes: 2})
	if err != nil {
		t.Fatalf("seed witness b: %v", err)
	}
	anchor, err := state.AddPassage(lexicon.Passage{WitnessID: a.ID, Location: lexicon.Location{Volume: 1, Leaf: 2, Line: 3, Character: 1}, Text: "山川异域 风月同天"})
	if err != nil {
		t.Fatalf("seed anchor passage: %v", err)
	}
	if _, err := state.AddPassage(lexicon.Passage{WitnessID: b.ID, Location: lexicon.Location{Volume: 1, Leaf: 2, Line: 3, Character: 1}, Text: "山河异域 风月同天"}); err != nil {
		t.Fatalf("seed other passage: %v", err)
	}

	collation, err := state.Start(anchor.ID, []lexicon.Key{a.ID, b.ID})
	if err != nil {
		t.Fatalf("start collation: %v", err)
	}
	task := state.Enqueue(collation.ID, 0)

	var wg sync.WaitGroup
	const overlap = 2
	wg.Add(overlap)
	for i := 0; i < overlap; i++ {
		go func() {
			defer wg.Done()
			_ = state.Run(task.ID)
		}()
	}
	wg.Wait()

	// 进度统计：尝试次数应与实际完成次数一致——交叠执行只完成一次。
	summary := state.Summary()
	if got := summary["attempts"].(int); got != 1 {
		t.Fatalf("attempts = %d, want 1 (one real execution)", got)
	}
	if got := summary["attempt_semantics"].(string); got != "single" {
		t.Fatalf("attempt_semantics = %q, want single", got)
	}
	if got := len(state.Divergences()); got != 1 {
		t.Fatalf("divergences = %d, want 1", got)
	}
	tasks := state.Tasks()
	if len(tasks) != 1 || tasks[0].Attempts != 1 || tasks[0].Status != "completed" {
		t.Fatalf("task state = %+v, want Attempts=1 Status=completed", tasks)
	}
}
