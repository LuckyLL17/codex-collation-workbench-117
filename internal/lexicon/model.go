package lexicon

import "time"

type Key string
type Witness struct {
	ID        Key       `json:"id"`
	Title     string    `json:"title"`
	Dynasty   string    `json:"dynasty"`
	Priority  int       `json:"priority"`
	Volumes   int       `json:"volumes"`
	CreatedAt time.Time `json:"created_at"`
}
type Location struct {
	Volume    int `json:"volume"`
	Leaf      int `json:"leaf"`
	Line      int `json:"line"`
	Character int `json:"character"`
}
type Passage struct {
	ID        Key      `json:"id"`
	WitnessID Key      `json:"witness_id"`
	Location  Location `json:"location"`
	Text      string   `json:"text"`
	Digest    string   `json:"digest"`
	Revision  int      `json:"revision"`
}
type Variant struct {
	ID        Key    `json:"id"`
	PassageID Key    `json:"passage_id"`
	Kind      string `json:"kind"`
	Reading   string `json:"reading"`
	Note      string `json:"note"`
}
type RuleSet struct {
	ID        Key                `json:"id"`
	Name      string             `json:"name"`
	Version   int                `json:"version"`
	Weights   map[string]float64 `json:"weights"`
	Published bool               `json:"published"`
}
type Candidate struct {
	WitnessID Key      `json:"witness_id"`
	Text      string   `json:"text"`
	Score     float64  `json:"score"`
	Reasons   []string `json:"reasons"`
}
type Divergence struct {
	ID         Key         `json:"id"`
	Anchor     Key         `json:"anchor"`
	Candidates []Candidate `json:"candidates"`
	State      string      `json:"state"`
	RuleDigest string      `json:"rule_digest"`
}
type Collation struct {
	ID        Key    `json:"id"`
	Anchor    Key    `json:"anchor"`
	Witnesses []Key  `json:"witnesses"`
	RuleSetID Key    `json:"rule_set_id"`
	Shards    int    `json:"shards"`
	Status    string `json:"status"`
}
type Task struct {
	ID          Key    `json:"id"`
	CollationID Key    `json:"collation_id"`
	Shard       int    `json:"shard"`
	Status      string `json:"status"`
	Attempts    int    `json:"attempts"`
}
