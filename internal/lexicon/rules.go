package lexicon

import "fmt"

func DefaultRules() RuleSet {
	return RuleSet{ID: "rules-v1", Name: "传统校勘基础权重", Version: 1, Weights: map[string]float64{"similarity": .65, "priority": .25, "same_volume": .1}, Published: true}
}
func ValidateRules(v RuleSet) error {
	if v.Name == "" {
		return fmt.Errorf("rule name required")
	}
	for _, key := range []string{"similarity", "priority", "same_volume"} {
		if v.Weights[key] < 0 {
			return fmt.Errorf("weight %s required", key)
		}
	}
	return nil
}
func RuleDigest(v RuleSet) string {
	return string(KeyOf("rules", v.Name, fmt.Sprint(v.Version), fmt.Sprint(v.Weights)))
}
