package game

import (
	"time"
)

func Loop(tickRate time.Duration, onTick func(float64)) {
	tick := time.Tick(time.Second / tickRate)
	start := time.Now().UnixNano()

	for range tick {
		now := time.Now().UnixNano()
		delta := float64(now-start) / float64(time.Second)
		start = now

		onTick(delta)
	}
}
