package ecs_test

import (
	"math"
	"testing"

	"github.com/robshape/deck/pkg/ecs"
)

func TestNewComponentManager(t *testing.T) {
	componentManager := ecs.NewComponentManager(ecs.MaxEntities)
	entitiesCount := componentManager.EntitiesCount()

	if componentManager == nil {
		t.Error("got nil, want non-nil")
	}
	if entitiesCount != ecs.MaxEntities {
		t.Errorf("got %d, want %d", entitiesCount, ecs.MaxEntities)
	}
}

func TestAddComponent(t *testing.T) {
	cases := []struct {
		name string
		in   []ecs.Component
		want ecs.Signature
	}{
		{"should add one component with mask", []ecs.Component{&testComponent1{}}, 1},
		{"should add many components with mask", []ecs.Component{&testComponent1{}, &testComponent2{}, &testComponent3{}}, 1 | 2 | 4},
		{"should not add duplicate component with mask", []ecs.Component{&testComponent1{}, &testComponent1{}, &testComponent1{}}, 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			componentManager := ecs.NewComponentManager(ecs.MaxEntities)
			entityManager := ecs.NewEntityManager(ecs.MaxEntities)
			entity, err := entityManager.CreateEntity()
			if err != nil {
				t.Error(err)
			}

			for _, component := range c.in {
				componentManager.AddComponent(entity, component)
			}
			componentsCount := componentManager.ComponentsCount(entity)
			signature := componentManager.Signature(entity)
			wantCount := int(math.Round(math.Sqrt(float64(c.want)))) // Count bits set in mask

			if componentsCount != wantCount {
				t.Errorf("got %d, want %d", componentsCount, wantCount)
			}
			if signature != c.want {
				t.Errorf("got %d, want %d", signature, c.want)
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
			entity, err := entityManager.CreateEntity()
			if err != nil {
				t.Error(err)
			}

			for _, component := range c.in {
				if err := componentManager.AddComponent(entity, component); err != nil {
					t.Error(err)
				}
			}
			componentsCount := componentManager.ComponentsCount(entity)
			inCount := len(c.in)
			signature := componentManager.Signature(entity)

			if componentsCount != inCount {
				t.Errorf("got %d, want %d", componentsCount, inCount)
			}
			if signature == 0 {
				t.Errorf("got %d, want non-zero", signature)
			}

			componentManager.DestroyComponents(entity)
			componentsCount = componentManager.ComponentsCount(entity)
			signature = componentManager.Signature(entity)

			if componentsCount != 0 {
				t.Errorf("got %d, want 0", componentsCount)
			}
			if signature != 0 {
				t.Errorf("got %d, want 0", signature)
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
			entity, err := entityManager.CreateEntity()
			if err != nil {
				t.Error(err)
			}

			for _, component := range c.in {
				if err := componentManager.AddComponent(entity, component); err != nil {
					t.Error(err)
				}
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
			entity, err := entityManager.CreateEntity()
			if err != nil {
				t.Error(err)
			}

			for _, component := range c.in {
				if err := componentManager.AddComponent(entity, component); err != nil {
					t.Error(err)
				}
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
			signature := componentManager.Signature(entity)

			if componentsCount != 0 {
				t.Errorf("got %d, want 0", componentsCount)
			}
			if signature != 0 {
				t.Errorf("got %d, want 0", signature)
			}
		})
	}
}

////////////////////////
// DATA, MOCKS, STUBS //
////////////////////////

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
