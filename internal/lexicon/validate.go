package lexicon

import "fmt"

func ValidateWitness(v Witness) error {
	if Normalize(v.Title) == "" {
		return fmt.Errorf("witness title required")
	}
	if v.Priority < 1 || v.Volumes < 1 {
		return fmt.Errorf("witness priority and volumes must be positive")
	}
	return nil
}
func ValidatePassage(v Passage) error {
	if !v.Location.Valid() {
		return fmt.Errorf("location incomplete")
	}
	if Normalize(v.Text) == "" {
		return fmt.Errorf("passage text required")
	}
	return nil
}
func ValidateVariant(v Variant) error {
	allowed := map[string]bool{"增衍": true, "脱文": true, "倒文": true, "通假": true, "形近": true}
	if !allowed[v.Kind] {
		return fmt.Errorf("unsupported variant kind")
	}
	return nil
}
