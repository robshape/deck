package ecs_test

import (
	"testing"

	"github.com/robshape/deck/pkg/ecs"
)

func TestNewSystemManager(t *testing.T) {
	systemManager := ecs.NewSystemManager()

	if systemManager == nil {
		t.Error("got nil, want non-nil")
	}
}

func TestRegisterSystem(t *testing.T) {
	cases := []struct {
		name string
		in   []ecs.System
	}{
		{"should register one system", []ecs.System{&testSystem{}}},
		{"should register many systems", []ecs.System{&testSystem{}, &testSystem{}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			systemManager := ecs.NewSystemManager()

			for _, system := range c.in {
				systemManager.RegisterSystem(system)
			}
			inCount := len(c.in)
			systemsCount := systemManager.SystemsCount()

			if systemsCount != inCount {
				t.Errorf("got %d, want %d", systemsCount, inCount)
			}
		})
	}
}

func TestUpdateEntities(t *testing.T) {
	cases := []struct {
		name string
		in   []ecs.System
	}{
		{"should update entities in one system", []ecs.System{&testSystem{}}},
		{"should update entities in many systems", []ecs.System{&testSystem{}, &testSystem{}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			componentManager := ecs.NewComponentManager(ecs.MaxEntities)
			entityManager := ecs.NewEntityManager(ecs.MaxEntities)
			systemManager := ecs.NewSystemManager()

			entity, err := entityManager.CreateEntity()
			if err != nil {
				t.Error(err)
			}
			componentManager.AddComponent(entity, &testSystemComponent{})

			for _, system := range c.in {
				systemManager.RegisterSystem(system)
				testSystem := system.(*testSystem)
				testSystemEntitiesCount := len(testSystem.entities)

				if testSystemEntitiesCount != 0 {
					t.Errorf("got %d, want 0", testSystemEntitiesCount)
				}
			}

			systemManager.UpdateEntities(entity, testSystemComponentSignature)

			for _, system := range c.in {
				testSystem := system.(*testSystem)
				testSystemEntitiesCount := len(testSystem.entities)

				if testSystemEntitiesCount != 1 {
					t.Errorf("got %d, want 1", testSystemEntitiesCount)
				}
			}

			systemManager.RemoveEntity(entity)

			for _, system := range c.in {
				testSystem := system.(*testSystem)
				testSystemEntitiesCount := len(testSystem.entities)

				if testSystemEntitiesCount != 0 {
					t.Errorf("got %d, want 0", testSystemEntitiesCount)
				}
			}
		})
	}
}

func TestUpdateSystems(t *testing.T) {
	type in struct {
		dt      float64
		systems []ecs.System
	}
	cases := []struct {
		name string
		in   in
	}{
		{"should update one system", in{60, []ecs.System{&testSystem{}}}},
		{"should update many systems", in{120, []ecs.System{&testSystem{}, &testSystem{}}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			systemManager := ecs.NewSystemManager()

			for _, system := range c.in.systems {
				systemManager.RegisterSystem(system)
				testSystem := system.(*testSystem)

				if testSystem.dt != 0 {
					t.Errorf("got %f, want 0", testSystem.dt)
				}
			}

			systemManager.UpdateSystems(c.in.dt)

			for _, system := range c.in.systems {
				testSystem := system.(*testSystem)

				if testSystem.dt != c.in.dt {
					t.Errorf("got %f, want %f", testSystem.dt, c.in.dt)
				}
			}
		})
	}
}

////////////////////////
// DATA, MOCKS, STUBS //
////////////////////////

const testSystemComponentSignature = 1

type testSystem struct {
	entities []ecs.Entity
	dt       float64
}

func (ts *testSystem) AddEntity(entity ecs.Entity) {
	ts.entities = append(ts.entities, entity)
}

func (ts *testSystem) RemoveEntity(entity ecs.Entity) {
	for i, e := range ts.entities {
		if e == entity {
			ts.entities[i] = ts.entities[len(ts.entities)-1]
			ts.entities[len(ts.entities)-1] = 0
			ts.entities = ts.entities[:len(ts.entities)-1]
			break
		}
	}
}

func (ts *testSystem) Signature() ecs.Signature {
	return testSystemComponentSignature
}

func (ts *testSystem) Update(dt float64) {
	ts.dt = dt
}

type testSystemComponent struct{}

func (tsc *testSystemComponent) Type() ecs.ComponentType {
	return testSystemComponentSignature
}
