package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

func CreateGameObjects(ecsManager *ecs.EcsManager) {
	createMarkers(ecsManager)
}
