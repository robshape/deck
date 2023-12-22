package game

import (
	"github.com/robshape/deck/pkg/ecs"
)

func Start() {
	ecsManager := ecs.NewEcsManager(ecs.MaxEntities)

	createGameObjects(ecsManager)

	const tickRate = 60 // Ticks per second
	loop(tickRate, func(dt float64) {
		// input()
		ecsManager.UpdateSystems(dt)
		// render()
	})
}
