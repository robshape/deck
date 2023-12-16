package ecs_test

import (
	"math"
	"testing"

	"github.com/robshape/deck/internal/ecs"
)

type testComponent1 struct{}

func (tc1 *testComponent1) Type() ecs.ComponentType {
	return 1
}

type testComponent2 struct{}

func (tc2 *testComponent2) Type() ecs.ComponentType {
	return 2
}

type testComponent3 struct{}

func (tc3 *testComponent3) Type() ecs.ComponentType {
	return 4
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
		in   []ecs.Component
		want ecs.ComponentsMask
	}{
		{"should add one component with mask", []ecs.Component{&testComponent1{}}, 1},
		{"should add many components with mask", []ecs.Component{&testComponent1{}, &testComponent2{}, &testComponent3{}}, 1 | 2 | 4},
		{"should not add duplicate component with mask", []ecs.Component{&testComponent1{}, &testComponent1{}, &testComponent1{}}, 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			componentManager := ecs.NewComponentManager(ecs.MaxEntities)
			entityManager := ecs.NewEntityManager(ecs.MaxEntities)
			entity, _ := entityManager.CreateEntity()

			for _, component := range c.in {
				componentManager.AddComponent(entity, component)
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
		in   []ecs.Component
	}{
		{"should destroy one component", []ecs.Component{&testComponent1{}}},
		{"should destroy many components", []ecs.Component{&testComponent1{}, &testComponent2{}, &testComponent3{}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			componentManager := ecs.NewComponentManager(ecs.MaxEntities)
			entityManager := ecs.NewEntityManager(ecs.MaxEntities)
			entity, _ := entityManager.CreateEntity()

			for _, component := range c.in {
				componentManager.AddComponent(entity, component)
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

func TestGetComponent(t *testing.T) {
	cases := []struct {
		name string
		in   []ecs.Component
	}{
		{"should get one component", []ecs.Component{&testComponent1{}}},
		{"should get many components", []ecs.Component{&testComponent1{}, &testComponent2{}, &testComponent3{}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			componentManager := ecs.NewComponentManager(ecs.MaxEntities)
			entityManager := ecs.NewEntityManager(ecs.MaxEntities)
			entity, _ := entityManager.CreateEntity()

			for _, component := range c.in {
				componentManager.AddComponent(entity, component)
			}
			componentsCount := componentManager.ComponentsCount(entity)
			inCount := len(c.in)

			if componentsCount != inCount {
				t.Errorf("got %d, want %d", componentsCount, inCount)
			}

			for _, component := range c.in {
				componentType := component.Type()
				component := componentManager.GetComponent(entity, componentType)

				if component == nil {
					t.Errorf("got nil, want non-nil")
				}
			}
		})
	}
}

func TestRemoveComponent(t *testing.T) {
	cases := []struct {
		name string
		in   []ecs.Component
	}{
		{"should remove one component", []ecs.Component{&testComponent1{}}},
		{"should remove many components", []ecs.Component{&testComponent1{}, &testComponent2{}, &testComponent3{}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			componentManager := ecs.NewComponentManager(ecs.MaxEntities)
			entityManager := ecs.NewEntityManager(ecs.MaxEntities)
			entity, _ := entityManager.CreateEntity()

			for _, component := range c.in {
				componentManager.AddComponent(entity, component)
			}
			componentsCount := componentManager.ComponentsCount(entity)
			inCount := len(c.in)

			if componentsCount != inCount {
				t.Errorf("got %d, want %d", componentsCount, inCount)
			}

			for _, component := range c.in {
				componentType := component.Type()
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
