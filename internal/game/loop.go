package game

import (
	"time"
)

func loop(tickRate time.Duration, onTick func(float64)) {
	tick := time.Tick(time.Second / tickRate)
	start := time.Now().UnixNano()

	for range tick {
		now := time.Now().UnixNano()
		dt := float64(now-start) / float64(time.Second)
		start = now

		onTick(dt)
	}
}
