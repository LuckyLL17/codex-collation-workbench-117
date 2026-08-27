package dispatch

type Metrics struct {
	Depth      int  `json:"depth"`
	Partitions int  `json:"partitions"`
	Closed     bool `json:"closed"`
}

func (m *Mailbox) Metrics() Metrics {
	depth := m.Depth()
	m.mu.Lock()
	partitions := m.partitions
	closed := m.closed
	m.mu.Unlock()
	return Metrics{Depth: depth, Partitions: partitions, Closed: closed}
}
