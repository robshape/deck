package ecs

import (
	"math"
	"testing"
)

const (
	COMPONENT_MOCK_1 = 1
	COMPONENT_MOCK_2 = 2
	COMPONENT_MOCK_3 = 4
)

type mockComponent struct {
	componentType ComponentType
}

func (mc *mockComponent) Type() ComponentType {
	return mc.componentType
}

func TestNewComponentManager(t *testing.T) {
	componentManager := newComponentManager(maxEntities)
	entityComponentsCount := len(componentManager.entityComponents)

	if componentManager == nil {
		t.Error("got nil, want non-nil")
	}
	if entityComponentsCount != maxEntities {
		t.Errorf("got %d, want %d", entityComponentsCount, maxEntities)
	}
}

func TestAddComponent(t *testing.T) {
	cases := []struct {
		name string
		in   []ComponentType
		want componentsMask
	}{
		{"should add one component with mask", []ComponentType{COMPONENT_MOCK_1}, COMPONENT_MOCK_1},
		{"should add many components with mask", []ComponentType{COMPONENT_MOCK_1, COMPONENT_MOCK_2, COMPONENT_MOCK_3}, COMPONENT_MOCK_1 | COMPONENT_MOCK_2 | COMPONENT_MOCK_3},
		{"should not add duplicate component with mask", []ComponentType{COMPONENT_MOCK_1, COMPONENT_MOCK_1, COMPONENT_MOCK_1}, COMPONENT_MOCK_1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			componentManager := newComponentManager(maxEntities)
			entityManager := newEntityManager(maxEntities)
			entity := entityManager.createEntity()

			for _, componentType := range c.in {
				componentManager.addComponent(entity, &mockComponent{componentType: componentType})
			}
			componentsCount := componentManager.componentsCount(entity)
			componentsMask := componentManager.componentsMask(entity)
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

func TestRemoveComponent(t *testing.T) {
	cases := []struct {
		name string
		in   []ComponentType
	}{
		{"should remove one component", []ComponentType{COMPONENT_MOCK_1}},
		{"should remove many components", []ComponentType{COMPONENT_MOCK_1, COMPONENT_MOCK_2, COMPONENT_MOCK_3}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			componentManager := newComponentManager(maxEntities)
			entityManager := newEntityManager(maxEntities)
			entity := entityManager.createEntity()

			for _, componentType := range c.in {
				componentManager.addComponent(entity, &mockComponent{componentType: componentType})
			}
			componentsCount := componentManager.componentsCount(entity)
			inCount := len(c.in)

			if componentsCount != inCount {
				t.Errorf("got %d, want %d", componentsCount, inCount)
			}

			for _, componentType := range c.in {
				componentManager.removeComponent(entity, componentType)
			}
			componentsCount = componentManager.componentsCount(entity)
			componentsMask := componentManager.componentsMask(entity)

			if componentsCount != 0 {
				t.Errorf("got %d, want 0", componentsCount)
			}
			if componentsMask != 0 {
				t.Errorf("got %d, want 0", componentsMask)
			}
		})
	}
}
