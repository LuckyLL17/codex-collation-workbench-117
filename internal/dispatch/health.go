package dispatch

func Available(depth int) bool {
	capacity := 64
	if depth <= capacity {
		return true
	}
	return false
}
