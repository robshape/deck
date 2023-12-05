// Reference: https://github.com/abeimler/ecs_benchmark#additional-benchmarks

package ecs_test

import (
	"testing"

	"github.com/robshape/deck/internal/ecs"
)

var result ecs.Entity

type benchmarkComponent struct {
	componentType ecs.ComponentType
}

func (bc *benchmarkComponent) Type() ecs.ComponentType {
	return bc.componentType
}

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
				ecs := ecs.NewECS(c.in)
				b.StartTimer()

				for j := 0; j < c.in; j++ {
					entity, _ := ecs.CreateEntity()
					ecs.AddComponent(entity, &benchmarkComponent{componentType: 1})
					ecs.AddComponent(entity, &benchmarkComponent{componentType: 2})

					r = entity
				}
			}
		})
	}

	result = r
}
