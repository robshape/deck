package game

import (
	"github.com/robshape/deck/internal/render"
	"github.com/robshape/deck/pkg/ecs"
	"github.com/robshape/deck/pkg/renderer"
)

func Start() {
	ecsManager := ecs.NewEcsManager(ecs.MaxEntities)

	renderer := renderer.NewRenderer()
	renderSystem := render.NewRenderSystem(renderer)
	ecsManager.RegisterSystem(renderSystem)

	createGameObjects(ecsManager)

	const tickRate = 60 // Ticks per second
	loop(tickRate, func(dt float64) {
		// input()
		ecsManager.UpdateSystems(dt)
	})
}
