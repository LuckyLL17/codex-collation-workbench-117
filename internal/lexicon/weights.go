package lexicon

func Clamp(value float64) float64 {
	if value < 0 {
		return 0
	}
	if value > 1 {
		return 1
	}
	return value
}
func Weighted(score, weight float64) float64 { return Clamp(score * weight) }
