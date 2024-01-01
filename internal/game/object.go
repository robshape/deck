package game

import (
	"fmt"
	"log"

	"github.com/robshape/deck/internal/game/component"
	"github.com/robshape/deck/internal/render"
)

func createGameObjects(ecsManager ecsManager) {
	createDamageCounters(ecsManager)
	createForceMarker(ecsManager)
	createResourceCounters(ecsManager)
}

func createDamageCounters(ecsManager ecsManager) {
	const damageCountersCount = 50
	for i := 0; i < damageCountersCount; i++ {
		damageCounterEntity, err := ecsManager.CreateEntity()
		if err != nil {
			log.Fatal(fmt.Errorf("createDamageCounters: failed creating damage counter entity: %w", err))
		}

		attackComponent := &component.AttackComponent{
			Damage: 1,
		}
		if err := ecsManager.AddComponent(damageCounterEntity, attackComponent); err != nil {
			log.Fatal(fmt.Errorf("createDamageCounters: failed adding attack component to damage counter entity: %w", err))
		}
	}
}

func createForceMarker(ecsManager ecsManager) {
	forceMarkerEntity, err := ecsManager.CreateEntity()
	if err != nil {
		log.Fatal(fmt.Errorf("createForceMarker: failed creating force marker entity: %w", err))
	}

	forceComponent := &component.ForceComponent{
		Force: 1,
	}
	if err := ecsManager.AddComponent(forceMarkerEntity, forceComponent); err != nil {
		log.Fatal(fmt.Errorf("createForceMarker: failed adding force component to force marker entity: %w", err))
	}

	renderComponent := &render.RenderComponent{}
	if err := ecsManager.AddComponent(forceMarkerEntity, renderComponent); err != nil {
		log.Fatal(fmt.Errorf("createForceMarker: failed adding render component to force marker entity: %w", err))
	}
}

func createResourceCounters(ecsManager ecsManager) {
	const resourceCountersCount = 20
	for i := 0; i < resourceCountersCount; i++ {
		resourceCounterEntity, err := ecsManager.CreateEntity()
		if err != nil {
			log.Fatal(fmt.Errorf("createResourceCounters: failed creating resource counter entity: %w", err))
		}

		resourcesComponent := &component.ResourcesComponent{
			Resources: 1,
		}
		if err := ecsManager.AddComponent(resourceCounterEntity, resourcesComponent); err != nil {
			log.Fatal(fmt.Errorf("createResourceCounters: failed adding resources component to resource counter entity: %w", err))
		}
	}
}
