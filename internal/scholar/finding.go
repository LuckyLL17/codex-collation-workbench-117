package scholar

type Finding struct {
	Rule        string  `json:"rule"`
	Score       float64 `json:"score"`
	Matched     bool    `json:"matched"`
	Explanation string  `json:"explanation"`
}
