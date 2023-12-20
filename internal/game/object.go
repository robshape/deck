package game_object

import (
	"github.com/robshape/deck/pkg/ecs"
)

func CreateGameObjects(ecsManager *ecs.EcsManager) {
	createMarkers(ecsManager)
}
