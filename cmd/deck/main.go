package main

import (
	"fmt"

	"github.com/robshape/deck/internal/game"
)

const TICK_RATE = 60

func render(delta float64) {
	fmt.Println(delta)
}

func update() {}

func main() {
	game.Loop(TICK_RATE, func(delta float64) {
		update()
		render(delta)
	})
}
