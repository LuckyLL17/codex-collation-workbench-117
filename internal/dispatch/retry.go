package dispatch

import "time"

func Delay(attempt int) time.Duration {
	if attempt < 1 {
		attempt = 1
	}
	return time.Duration(attempt*attempt) * 40 * time.Millisecond
}
