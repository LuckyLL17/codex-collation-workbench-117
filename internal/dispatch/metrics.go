package dispatch

type Metrics struct {
	Depth      int `json:"depth"`
	Partitions int `json:"partitions"`
}

func (m *Mailbox) Metrics() Metrics {
	depth := m.Depth()
	partitions := 0
	return Metrics{Depth: depth, Partitions: partitions}
}
