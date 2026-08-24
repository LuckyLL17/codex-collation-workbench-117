package lexicon

func Penalty(kind string) float64 {
	switch kind {
	case "脱文":
		return .35
	case "增衍":
		return .2
	case "倒文":
		return .18
	case "通假":
		return .08
	default:
		return .12
	}
}
func Label(kind string) string {
	labels := map[string]string{"脱文": "疑似抄写遗漏", "增衍": "疑似后人补入", "倒文": "疑似次序移动", "通假": "语词替换", "形近": "字形相近"}
	return labels[kind]
}
