package game

import (
	"github.com/robshape/deck/internal/component"
	"github.com/robshape/deck/pkg/ecs"
)

func createGameObjects(ecsManager *ecs.EcsManager) {
	createMarkers(ecsManager)
}

func createMarkers(ecsManager *ecs.EcsManager) {
	const damageCountersCount = 50
	for i := 0; i < damageCountersCount; i++ {
		damageCounterEntity, _ := ecsManager.CreateEntity()
		ecsManager.AddComponent(damageCounterEntity, &component.AttackComponent{
			Damage: 1,
		})
	}

	forceMarkerEntity, _ := ecsManager.CreateEntity()
	ecsManager.AddComponent(forceMarkerEntity, &component.ForceComponent{
		Force: 1,
	})

	const resourceCountersCount = 20
	for i := 0; i < resourceCountersCount; i++ {
		resourceCounterEntity, _ := ecsManager.CreateEntity()
		ecsManager.AddComponent(resourceCounterEntity, &component.ResourcesComponent{
			Resources: 1,
		})
	}
}
