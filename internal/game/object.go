package game

import (
	"fmt"
	"log"

	"github.com/robshape/deck/internal/game/component"
	"github.com/robshape/deck/pkg/ecs"
)

func createGameObjects(ecsManager ecs.EcsManager) {
	createMarkers(ecsManager)
}

func createMarkers(ecsManager ecs.EcsManager) {
	const damageCountersCount = 50
	for i := 0; i < damageCountersCount; i++ {
		damageCounterEntity, err := ecsManager.CreateEntity()
		if err != nil {
			log.Fatal(fmt.Errorf("createMarkers: failed creating damage counter: %w", err))
		}

		attackComponent := &component.AttackComponent{
			Damage: 1,
		}
		if err := ecsManager.AddComponent(damageCounterEntity, attackComponent); err != nil {
			log.Fatal(fmt.Errorf("createMarkers: failed adding attack to damage counter: %w", err))
		}
	}

	forceMarkerEntity, err := ecsManager.CreateEntity()
	if err != nil {
		log.Fatal(fmt.Errorf("createMarkers: failed creating force marker: %w", err))
	}

	forceComponent := &component.ForceComponent{
		Force: 1,
	}
	if err := ecsManager.AddComponent(forceMarkerEntity, forceComponent); err != nil {
		log.Fatal(fmt.Errorf("createMarkers: failed adding force to force marker: %w", err))
	}

	const resourceCountersCount = 20
	for i := 0; i < resourceCountersCount; i++ {
		resourceCounterEntity, err := ecsManager.CreateEntity()
		if err != nil {
			log.Fatal(fmt.Errorf("createMarkers: failed creating resource counter: %w", err))
		}

		resourcesComponent := &component.ResourcesComponent{
			Resources: 1,
		}
		if err := ecsManager.AddComponent(resourceCounterEntity, resourcesComponent); err != nil {
			log.Fatal(fmt.Errorf("createMarkers: failed adding resources to resource counter: %w", err))
		}
	}
}
