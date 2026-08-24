package lexicon

import "fmt"

func PassageLabel(v Passage) string {
	return fmt.Sprintf("%s @ %s", v.Digest[:12], v.Location.String())
}
func WitnessLabel(v Witness) string { return fmt.Sprintf("%s（%s）", v.Title, v.Dynasty) }
