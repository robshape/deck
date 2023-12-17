package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

const (
	// Automatically increment by the power of two
	descriptionComponentType ecs.ComponentType = 1 << iota
	healthComponentType
	propertiesComponentType
)

func CreateGameObjects(ecsManager *ecs.EcsManager) {
	const damageCountersCount = 50
	for i := 0; i < damageCountersCount; i++ {
		damageCounterEntity, _ := ecsManager.CreateEntity()
		ecsManager.AddComponent(damageCounterEntity, &propertiesComponent{
			attack: 1,
		})
	}

	forceMarkerEntity, _ := ecsManager.CreateEntity()
	ecsManager.AddComponent(forceMarkerEntity, &propertiesComponent{
		force: 1,
	})

	const resourceCountersCount = 20
	for i := 0; i < resourceCountersCount; i++ {
		resourceCounterEntity, _ := ecsManager.CreateEntity()
		ecsManager.AddComponent(resourceCounterEntity, &propertiesComponent{
			resources: 1,
		})
	}
}
