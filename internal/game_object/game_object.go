package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

const (
	// Automatically increment by the power of two
	abilityComponentType ecs.ComponentType = 1 << iota
	attackComponentType
	cardComponentType
	costComponentType
	factionComponentType
	forceComponentType
	healthComponentType
	resourcesComponentType
	traitComponentType
)

func CreateGameObjects(ecsManager *ecs.EcsManager) {
	const damageCountersCount = 50
	for i := 0; i < damageCountersCount; i++ {
		damageCounterEntity, _ := ecsManager.CreateEntity()
		ecsManager.AddComponent(damageCounterEntity, &attackComponent{
			damage: 1,
		})
	}

	forceMarkerEntity, _ := ecsManager.CreateEntity()
	ecsManager.AddComponent(forceMarkerEntity, &forceComponent{
		force: 1,
	})

	const resourceCountersCount = 20
	for i := 0; i < resourceCountersCount; i++ {
		resourceCounterEntity, _ := ecsManager.CreateEntity()
		ecsManager.AddComponent(resourceCounterEntity, &resourcesComponent{
			resources: 1,
		})
	}
}
