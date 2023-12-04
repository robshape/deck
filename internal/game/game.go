package game

import (
	"fmt"

	"github.com/robshape/deck/internal/ecs"
)

const (
	// Automatically increment by the power of two
	COMPONENT_DAMAGE_COUNTER ecs.ComponentType = 1 << iota
	COMPONENT_FORCE_MARKER
	COMPONENT_RESOURCE_COUNTER
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
