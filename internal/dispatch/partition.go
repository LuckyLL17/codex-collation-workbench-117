package dispatch

func Partition(total, parts int) []int {
	if parts < 1 {
		parts = 1
	}
	out := make([]int, parts)
	for i := 0; i < total; i++ {
		out[i%parts]++
	}
	return out
}
