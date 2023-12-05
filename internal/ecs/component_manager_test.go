package ecs_test

import (
	"math"
	"testing"

	"github.com/robshape/deck/internal/ecs"
)

const (
	componentMock1 = 1
	componentMock2 = 2
	componentMock3 = 4
)

type testComponent struct {
	componentType ecs.ComponentType
}

func (mc *testComponent) Type() ecs.ComponentType {
	return mc.componentType
}

func TestNewComponentManager(t *testing.T) {
	componentManager := ecs.NewComponentManager(ecs.MaxEntities)
	entityComponentsCount := componentManager.EntityComponentsCount()

	if componentManager == nil {
		t.Error("got nil, want non-nil")
	}
	if entityComponentsCount != ecs.MaxEntities {
		t.Errorf("got %d, want %d", entityComponentsCount, ecs.MaxEntities)
	}
}

func TestAddComponent(t *testing.T) {
	cases := []struct {
		name string
		in   []ecs.ComponentType
		want ecs.ComponentsMask
	}{
		{"should add one component with mask", []ecs.ComponentType{componentMock1}, componentMock1},
		{"should add many components with mask", []ecs.ComponentType{componentMock1, componentMock2, componentMock3}, componentMock1 | componentMock2 | componentMock3},
		{"should not add duplicate component with mask", []ecs.ComponentType{componentMock1, componentMock1, componentMock1}, componentMock1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			componentManager := ecs.NewComponentManager(ecs.MaxEntities)
			entityManager := ecs.NewEntityManager(ecs.MaxEntities)
			entity, _ := entityManager.CreateEntity()

			for _, componentType := range c.in {
				componentManager.AddComponent(entity, &testComponent{componentType: componentType})
			}
			componentsCount := componentManager.ComponentsCount(entity)
			componentsMask := componentManager.ComponentsMask(entity)
			wantCount := int(math.Round(math.Sqrt(float64(c.want)))) // Count bits set in mask

			if componentsCount != wantCount {
				t.Errorf("got %d, want %d", componentsCount, wantCount)
			}
			if componentsMask != c.want {
				t.Errorf("got %d, want %d", componentsMask, c.want)
			}
		})
	}
}

func TestDestroyEntityComponents(t *testing.T) {
	cases := []struct {
		name string
		in   []ecs.ComponentType
	}{
		{"should destroy one component", []ecs.ComponentType{componentMock1}},
		{"should destroy many components", []ecs.ComponentType{componentMock1, componentMock2, componentMock3}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			componentManager := ecs.NewComponentManager(ecs.MaxEntities)
			entityManager := ecs.NewEntityManager(ecs.MaxEntities)
			entity, _ := entityManager.CreateEntity()

			for _, componentType := range c.in {
				componentManager.AddComponent(entity, &testComponent{componentType: componentType})
			}
			componentsCount := componentManager.ComponentsCount(entity)
			componentsMask := componentManager.ComponentsMask(entity)
			inCount := len(c.in)

			if componentsCount != inCount {
				t.Errorf("got %d, want %d", componentsCount, inCount)
			}
			if componentsMask == 0 {
				t.Errorf("got %d, want non-zero", componentsMask)
			}

			componentManager.DestroyEntityComponents(entity)
			componentsCount = componentManager.ComponentsCount(entity)
			componentsMask = componentManager.ComponentsMask(entity)

			if componentsCount != 0 {
				t.Errorf("got %d, want 0", componentsCount)
			}
			if componentsMask != 0 {
				t.Errorf("got %d, want 0", componentsMask)
			}
		})
	}
}

func TestRemoveComponent(t *testing.T) {
	cases := []struct {
		name string
		in   []ecs.ComponentType
	}{
		{"should remove one component", []ecs.ComponentType{componentMock1}},
		{"should remove many components", []ecs.ComponentType{componentMock1, componentMock2, componentMock3}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			componentManager := ecs.NewComponentManager(ecs.MaxEntities)
			entityManager := ecs.NewEntityManager(ecs.MaxEntities)
			entity, _ := entityManager.CreateEntity()

			for _, componentType := range c.in {
				componentManager.AddComponent(entity, &testComponent{componentType: componentType})
			}
			componentsCount := componentManager.ComponentsCount(entity)
			inCount := len(c.in)

			if componentsCount != inCount {
				t.Errorf("got %d, want %d", componentsCount, inCount)
			}

			for _, componentType := range c.in {
				componentManager.RemoveComponent(entity, componentType)
			}
			componentsCount = componentManager.ComponentsCount(entity)
			componentsMask := componentManager.ComponentsMask(entity)

			if componentsCount != 0 {
				t.Errorf("got %d, want 0", componentsCount)
			}
			if componentsMask != 0 {
				t.Errorf("got %d, want 0", componentsMask)
			}
		})
	}
}
