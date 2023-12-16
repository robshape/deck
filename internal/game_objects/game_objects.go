package game_objects

import (
	"github.com/robshape/deck/internal/ecs"
)

const (
	// Automatically increment by the power of two
	damageCounterComponentType ecs.ComponentType = 1 << iota
	forceMarkerComponentType
	resourceCounterComponentType
)

func CreateGameObjects(ecsManager *ecs.EcsManager) {
	const damageCountersCount = 50
	for i := 0; i < damageCountersCount; i++ {
		damageCounterEntity, _ := ecsManager.CreateEntity()
		ecsManager.AddComponent(damageCounterEntity, &DamageCounterComponent{})
	}

	forceMarkerEntity, _ := ecsManager.CreateEntity()
	ecsManager.AddComponent(forceMarkerEntity, &ForceMarkerComponent{})

	const resourceCountersCount = 20
	for i := 0; i < resourceCountersCount; i++ {
		resourceCounterEntity, _ := ecsManager.CreateEntity()
		ecsManager.AddComponent(resourceCounterEntity, &ResourceCounterComponent{})
	}
}
