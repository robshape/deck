package game

import (
	"fmt"
)

func input() {}

func render(delta float64) {
	fmt.Println(delta)
}

func update() {}

func Start() {
	const TICK_RATE = 60 // Ticks per second
	loop(TICK_RATE, func(delta float64) {
		input()
		update()
		render(delta)
	})
}
