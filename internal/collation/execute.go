package collation

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func (s *State) Run(id lexicon.Key) error {
	s.mu.Lock()
	task, ok := s.tasks[id]
	if !ok {
		s.mu.Unlock()
		return ErrMissing
	}
	// 后台执行可能在交叠时对同一任务重复触发：已进入 running 或已 completed
	// 的任务表示执行体已接管，重复调用必须直接退出，既不重复执行、也不重复计数，
	// 否则进度统计的尝试次数会高于实际完成次数，监控会误判为发生重试。
	if task.Status == "running" || task.Status == "completed" {
		s.mu.Unlock()
		return nil
	}
	current := task.Attempts
	task.Attempts = current + 1
	task.Status = "running"
	s.tasks[id] = task
	batch := s.collations[task.CollationID]
	base := s.passages[batch.Anchor]
	others := []lexicon.Passage{}
	for _, value := range s.passages {
		if value.ID != base.ID {
			others = append(others, value)
		}
	}
	rules := s.rules
	s.mu.Unlock()
	candidates := []lexicon.Candidate{}
	for _, value := range others {
		candidates = append(candidates, lexicon.Score(base, value, s.witnesses[value.WitnessID], rules))
	}
	s.mu.Lock()
	divergenceID := lexicon.KeyOf("divergence", string(id))
	s.divergences[divergenceID] = lexicon.Divergence{ID: divergenceID, Anchor: base.ID, Candidates: candidates, State: "needs-review", RuleDigest: lexicon.RuleDigest(rules)}
	task.Status = "completed"
	s.tasks[id] = task
	s.mu.Unlock()
	return nil
}
func (s *State) Divergences() []lexicon.Divergence {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []lexicon.Divergence{}
	for _, value := range s.divergences {
		out = append(out, value)
	}
	return out
}
