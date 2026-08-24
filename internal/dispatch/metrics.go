package dispatch

type Metrics struct {
	Depth      int `json:"depth"`
	Partitions int `json:"partitions"`
}

func (m *Mailbox) Metrics() Metrics { return Metrics{Depth: m.Depth(), Partitions: 1} }
