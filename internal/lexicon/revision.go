package lexicon

func NextRevision(value Passage) Passage {
	value.Revision++
	value.Digest = RevisionDigest(value.Text)
	return value
}
func SameCoordinate(a, b Passage) bool {
	return a.Location == b.Location
}
