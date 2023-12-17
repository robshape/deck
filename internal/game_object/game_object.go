package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

const (
	// Automatically increment by the power of two
	cardComponentType ecs.ComponentType = 1 << iota
	markerComponentType
	propertiesComponentType
)

func CreateGameObjects(ecsManager *ecs.EcsManager) {
	const damageCountersCount = 50
	for i := 0; i < damageCountersCount; i++ {
		damageCounterEntity, _ := ecsManager.CreateEntity()
		ecsManager.AddComponent(damageCounterEntity, &MarkerComponent{MarkerTypeDamageCounter})
	}

	forceMarkerEntity, _ := ecsManager.CreateEntity()
	ecsManager.AddComponent(forceMarkerEntity, &MarkerComponent{MarkerTypeForceMarker})

	const resourceCountersCount = 20
	for i := 0; i < resourceCountersCount; i++ {
		resourceCounterEntity, _ := ecsManager.CreateEntity()
		ecsManager.AddComponent(resourceCounterEntity, &MarkerComponent{MarkerTypeResourceCounter})
	}
}
