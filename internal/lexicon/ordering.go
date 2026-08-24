package lexicon

func Prefer(a, b Witness) Witness {
	if a.Priority >= b.Priority {
		return a
	}
	return b
}
func SortKey(v Passage) string { return v.WitnessID.String() + v.Location.String() }
func (k Key) String() string   { return string(k) }
