// Reference: https://github.com/abeimler/ecs_benchmark#additional-benchmarks

package ecs_test

import (
	"testing"

	"github.com/robshape/deck/internal/ecs"
)

const (
	componentBenchmark1 = 1
	componentBenchmark2 = 2
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
				ecsCoordinator := ecs.NewECSCoordinator(c.in)
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
				ecsCoordinator := ecs.NewECSCoordinator(c.in)
				entities := []ecs.Entity{}
				for j := 0; j < c.in; j++ {
					entity, _ := ecsCoordinator.CreateEntity()
					ecsCoordinator.AddComponent(entity, &benchmarkComponent{componentType: componentBenchmark1})
					ecsCoordinator.AddComponent(entity, &benchmarkComponent{componentType: componentBenchmark2})
					entities = append(entities, entity)
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

// func BenchmarkGetComponent(b *testing.B) {}

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
				ecsCoordinator := ecs.NewECSCoordinator(c.in)
				entities := []ecs.Entity{}
				for j := 0; j < c.in; j++ {
					entity, _ := ecsCoordinator.CreateEntity()
					ecsCoordinator.AddComponent(entity, &benchmarkComponent{componentType: componentBenchmark1})
					entities = append(entities, entity)
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
