package game

import (
	"github.com/robshape/deck/internal/ecs"
	"github.com/robshape/deck/internal/game_objects"
)

func Start() {
	const tickRate = 60 // Ticks per second

	ecsManager := ecs.NewECSManager(ecs.MaxEntities)
	game_objects.CreateGameObjects(ecsManager)

	loop(tickRate, func(dt float64) {
		// input()
		ecsManager.UpdateSystems(dt)
		// render()
	})
}
