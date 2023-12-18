package game_object

import (
	gc "github.com/robshape/deck/internal/game_component"
	"github.com/robshape/deck/pkg/ecs"
)

func createMarkers(ecsManager *ecs.EcsManager) {
	const damageCountersCount = 50
	for i := 0; i < damageCountersCount; i++ {
		damageCounterEntity, _ := ecsManager.CreateEntity()
		ecsManager.AddComponent(damageCounterEntity, &gc.AttackComponent{
			Damage: 1,
		})
	}

	forceMarkerEntity, _ := ecsManager.CreateEntity()
	ecsManager.AddComponent(forceMarkerEntity, &gc.ForceComponent{
		Force: 1,
	})

	const resourceCountersCount = 20
	for i := 0; i < resourceCountersCount; i++ {
		resourceCounterEntity, _ := ecsManager.CreateEntity()
		ecsManager.AddComponent(resourceCounterEntity, &gc.ResourcesComponent{
			Resources: 1,
		})
	}
}
