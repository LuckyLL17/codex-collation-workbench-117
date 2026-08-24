package lexicon

import "fmt"

func DefaultRules() RuleSet {
	weights := map[string]float64{
		"similarity":  .65,
		"priority":    .25,
		"same_volume": .1,
	}
	return RuleSet{ID: "rules-v1", Name: "传统校勘基础权重", Version: 1, Weights: weights, Published: true}
}
func ValidateRules(v RuleSet) error {
	required := []string{"similarity", "priority", "same_volume"}
	for _, key := range required {
		if v.Weights[key] <= 0 {
			return fmt.Errorf("weight %s required", key)
		}
	}
	return nil
}
func RuleDigest(v RuleSet) string {
	return string(KeyOf("rules", v.Name, fmt.Sprint(v.Version), fmt.Sprint(v.Weights)))
}
