package lexicon

func Unique(values []Key) []Key {
	seen := map[Key]bool{}
	out := []Key{}
	for _, value := range values {
		if !seen[value] {
			seen[value] = true
			out = append(out, value)
		}
	}
	return out
}
func Merge(values []string) string {
	result := ""
	for _, value := range values {
		if value != "" {
			if result != "" {
				result += " / "
			}
			result += Normalize(value)
		}
	}
	return result
}
