package game

import (
	"time"
)

type onTickFn func(dt float64)

func loop(tickRate time.Duration, onTick onTickFn) {
	tick := time.Tick(time.Second / tickRate)
	start := time.Now().UnixNano()

	for range tick {
		now := time.Now().UnixNano()
		dt := float64(now-start) / float64(time.Second)
		start = now

		onTick(dt)
	}
}
