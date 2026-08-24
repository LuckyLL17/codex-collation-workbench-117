package dispatch

import "time"

func Tick() time.Time { return time.Now().UTC() }
