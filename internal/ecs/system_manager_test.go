package ecs_test

import (
	"testing"

	"github.com/robshape/deck/internal/ecs"
)

type testSystem struct {
	dt float64
}

func (ts *testSystem) Update(dt float64) {
	ts.dt = dt
}

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
