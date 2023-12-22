// Reference: https://github.com/abeimler/ecs_benchmark#additional-benchmarks

package ecs_test

import (
	"testing"

	"github.com/robshape/deck/pkg/ecs"
)

type benchmarkComponent1 struct{}

func (bc1 *benchmarkComponent1) Type() ecs.ComponentType {
	return 1
}

type benchmarkComponent2 struct{}

func (bc2 *benchmarkComponent2) Type() ecs.ComponentType {
	return 2
}

type benchmarkComponent3 struct{}

func (bc3 *benchmarkComponent3) Type() ecs.ComponentType {
	return 4
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
				ecsManager := ecs.NewEcsManager(c.in)
				b.StartTimer()

				for j := 0; j < c.in; j++ {
					entity, _ := ecsManager.CreateEntity()
					ecsManager.AddComponent(entity, &benchmarkComponent1{})
					ecsManager.AddComponent(entity, &benchmarkComponent2{})

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
				ecsManager := ecs.NewEcsManager(c.in)
				entities := []ecs.Entity{}
				for j := 0; j < c.in; j++ {
					entity, _ := ecsManager.CreateEntity()
					entities = append(entities, entity)
					ecsManager.AddComponent(entity, &benchmarkComponent1{})
					ecsManager.AddComponent(entity, &benchmarkComponent2{})
				}
				b.StartTimer()

				for j := 0; j < c.in; j++ {
					entity := entities[j]
					ecsManager.DestroyEntity(entity)

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
		entities   int
		components []ecs.Component
	}
	cases := []struct {
		name string
		in   in
	}{
		{"16 entities, 1 components", in{16, []ecs.Component{&benchmarkComponent1{}}}},
		{"64 entities, 1 components", in{64, []ecs.Component{&benchmarkComponent1{}}}},
		{"256 entities, 1 components", in{256, []ecs.Component{&benchmarkComponent1{}}}},
		{"1024 entities, 1 components", in{1024, []ecs.Component{&benchmarkComponent1{}}}},
		{"4096 entities, 1 components", in{4096, []ecs.Component{&benchmarkComponent1{}}}},

		{"16 entities, 2 components", in{16, []ecs.Component{&benchmarkComponent1{}, &benchmarkComponent2{}}}},
		{"64 entities, 2 components", in{64, []ecs.Component{&benchmarkComponent1{}, &benchmarkComponent2{}}}},
		{"256 entities, 2 components", in{256, []ecs.Component{&benchmarkComponent1{}, &benchmarkComponent2{}}}},
		{"1024 entities, 2 components", in{1024, []ecs.Component{&benchmarkComponent1{}, &benchmarkComponent2{}}}},
		{"4096 entities, 2 components", in{4096, []ecs.Component{&benchmarkComponent1{}, &benchmarkComponent2{}}}},

		{"16 entities, 3 components", in{16, []ecs.Component{&benchmarkComponent1{}, &benchmarkComponent2{}, &benchmarkComponent3{}}}},
		{"64 entities, 3 components", in{64, []ecs.Component{&benchmarkComponent1{}, &benchmarkComponent2{}, &benchmarkComponent3{}}}},
		{"256 entities, 3 components", in{256, []ecs.Component{&benchmarkComponent1{}, &benchmarkComponent2{}, &benchmarkComponent3{}}}},
		{"1024 entities, 3 components", in{1024, []ecs.Component{&benchmarkComponent1{}, &benchmarkComponent2{}, &benchmarkComponent3{}}}},
		{"4096 entities, 3 components", in{4096, []ecs.Component{&benchmarkComponent1{}, &benchmarkComponent2{}, &benchmarkComponent3{}}}},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				ecsManager := ecs.NewEcsManager(c.in.entities)
				entities := []ecs.Entity{}
				for j := 0; j < c.in.entities; j++ {
					entity, _ := ecsManager.CreateEntity()
					entities = append(entities, entity)

					for _, component := range c.in.components {
						ecsManager.AddComponent(entity, component)
					}
				}
				b.StartTimer()

				for j := 0; j < c.in.entities; j++ {
					entity := entities[j]

					for _, component := range c.in.components {
						componentType := component.Type()
						ecsManager.GetComponent(entity, componentType)
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
				benchmarkComponent1Type := (&benchmarkComponent1{}).Type()
				ecsManager := ecs.NewEcsManager(c.in)
				entities := []ecs.Entity{}
				for j := 0; j < c.in; j++ {
					entity, _ := ecsManager.CreateEntity()
					entities = append(entities, entity)
					ecsManager.AddComponent(entity, &benchmarkComponent1{})
				}
				b.StartTimer()

				for j := 0; j < c.in; j++ {
					entity := entities[j]
					ecsManager.RemoveComponent(entity, benchmarkComponent1Type)
					ecsManager.AddComponent(entity, &benchmarkComponent2{})

					r = entity
				}
			}
		})
	}

	result = r
}
