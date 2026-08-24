package shelf

import "github.com/local/codex-collation-workbench-117/internal/lexicon"

func Lineage(values []lexicon.Passage) map[string][]string {
	result := map[string][]string{}
	for _, value := range values {
		result[string(value.WitnessID)] = append(result[string(value.WitnessID)], value.Digest)
	}
	return result
}
