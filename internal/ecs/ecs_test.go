// Reference: https://github.com/abeimler/ecs_benchmark#additional-benchmarks

package ecs_test

import (
	"testing"

	"github.com/robshape/deck/internal/ecs"
)

const (
	componentBenchmark1 = 1
	componentBenchmark2 = 2
	componentBenchmark3 = 4
)

type benchmarkComponent struct {
	componentType ecs.ComponentType
}

func (bc *benchmarkComponent) Type() ecs.ComponentType {
	return bc.componentType
}

var result ecs.Entity

func BenchmarkCreateEntities(b *testing.B) {
	var r ecs.Entity // Avoid inline optimization (https://100go.co/89-benchmarks/)

	cases := []struct {
		name string
		in   int
	}{
		{"16 entities, 2 components", 16},
		{"64 entities, 2 components", 64},
		{"256 entities, 2 components", 256},
		{"1024 entities, 2 components", 1024},
		{"4096 entities, 2 components", 4096},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Avoid observer effect (https://100go.co/89-benchmarks/)
				b.StopTimer()
				ecsCoordinator := ecs.NewECSManager(c.in)
				b.StartTimer()

				for j := 0; j < c.in; j++ {
					entity, _ := ecsCoordinator.CreateEntity()
					ecsCoordinator.AddComponent(entity, &benchmarkComponent{componentType: componentBenchmark1})
					ecsCoordinator.AddComponent(entity, &benchmarkComponent{componentType: componentBenchmark2})

					r = entity
				}
			}
		})
	}

	result = r
}

func BenchmarkDestroyEntities(b *testing.B) {
	var r ecs.Entity

	cases := []struct {
		name string
		in   int
	}{
		{"16 entities, 2 components", 16},
		{"64 entities, 2 components", 64},
		{"256 entities, 2 components", 256},
		{"1024 entities, 2 components", 1024},
		{"4096 entities, 2 components", 4096},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				ecsCoordinator := ecs.NewECSManager(c.in)
				entities := []ecs.Entity{}
				for j := 0; j < c.in; j++ {
					entity, _ := ecsCoordinator.CreateEntity()
					entities = append(entities, entity)
					ecsCoordinator.AddComponent(entity, &benchmarkComponent{componentType: componentBenchmark1})
					ecsCoordinator.AddComponent(entity, &benchmarkComponent{componentType: componentBenchmark2})
				}
				b.StartTimer()

				for j := 0; j < c.in; j++ {
					entity := entities[j]
					ecsCoordinator.DestroyEntity(entity)

					r = entity
				}
			}
		})
	}

	result = r
}

func BenchmarkGetComponent(b *testing.B) {
	var r ecs.Entity

	type in struct {
		entities       int
		componentTypes []ecs.ComponentType
	}
	cases := []struct {
		name string
		in   in
	}{
		{"16 entities, 1 components", in{16, []ecs.ComponentType{componentBenchmark1}}},
		{"64 entities, 1 components", in{64, []ecs.ComponentType{componentBenchmark1}}},
		{"256 entities, 1 components", in{256, []ecs.ComponentType{componentBenchmark1}}},
		{"1024 entities, 1 components", in{1024, []ecs.ComponentType{componentBenchmark1}}},
		{"4096 entities, 1 components", in{4096, []ecs.ComponentType{componentBenchmark1}}},

		{"16 entities, 2 components", in{16, []ecs.ComponentType{componentBenchmark1, componentBenchmark2}}},
		{"64 entities, 2 components", in{64, []ecs.ComponentType{componentBenchmark1, componentBenchmark2}}},
		{"256 entities, 2 components", in{256, []ecs.ComponentType{componentBenchmark1, componentBenchmark2}}},
		{"1024 entities, 2 components", in{1024, []ecs.ComponentType{componentBenchmark1, componentBenchmark2}}},
		{"4096 entities, 2 components", in{4096, []ecs.ComponentType{componentBenchmark1, componentBenchmark2}}},

		{"16 entities, 3 components", in{16, []ecs.ComponentType{componentBenchmark1, componentBenchmark2, componentBenchmark3}}},
		{"64 entities, 3 components", in{64, []ecs.ComponentType{componentBenchmark1, componentBenchmark2, componentBenchmark3}}},
		{"256 entities, 3 components", in{256, []ecs.ComponentType{componentBenchmark1, componentBenchmark2, componentBenchmark3}}},
		{"1024 entities, 3 components", in{1024, []ecs.ComponentType{componentBenchmark1, componentBenchmark2, componentBenchmark3}}},
		{"4096 entities, 3 components", in{4096, []ecs.ComponentType{componentBenchmark1, componentBenchmark2, componentBenchmark3}}},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				ecsCoordinator := ecs.NewECSManager(c.in.entities)
				entities := []ecs.Entity{}
				for j := 0; j < c.in.entities; j++ {
					entity, _ := ecsCoordinator.CreateEntity()
					entities = append(entities, entity)

					for _, componentType := range c.in.componentTypes {
						ecsCoordinator.AddComponent(entity, &benchmarkComponent{componentType: componentType})
					}
				}
				b.StartTimer()

				for j := 0; j < c.in.entities; j++ {
					entity := entities[j]

					for _, componentType := range c.in.componentTypes {
						ecsCoordinator.GetComponent(entity, componentType)
					}

					r = entity
				}
			}
		})
	}

	result = r
}

func BenchmarkRemoveAddComponent(b *testing.B) {
	var r ecs.Entity

	cases := []struct {
		name string
		in   int
	}{
		{"16 entities, 1 component", 16},
		{"64 entities, 1 component", 64},
		{"256 entities, 1 component", 256},
		{"1024 entities, 1 component", 1024},
		{"4096 entities, 1 component", 4096},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				ecsCoordinator := ecs.NewECSManager(c.in)
				entities := []ecs.Entity{}
				for j := 0; j < c.in; j++ {
					entity, _ := ecsCoordinator.CreateEntity()
					entities = append(entities, entity)
					ecsCoordinator.AddComponent(entity, &benchmarkComponent{componentType: componentBenchmark1})
				}
				b.StartTimer()

				for j := 0; j < c.in; j++ {
					entity := entities[j]
					ecsCoordinator.RemoveComponent(entity, componentBenchmark1)
					ecsCoordinator.AddComponent(entity, &benchmarkComponent{componentType: componentBenchmark2})

					r = entity
				}
			}
		})
	}

	result = r
}
