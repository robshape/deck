package ecs_test

import (
	"testing"

	"github.com/robshape/deck/internal/ecs"
)

func TestNewEntityManager(t *testing.T) {
	entityManager := ecs.NewEntityManager(ecs.MaxEntities)
	_, destroyedEntitiesCount := entityManager.EntitiesCount()

	if entityManager == nil {
		t.Error("got nil, want non-nil")
	}
	if destroyedEntitiesCount != ecs.MaxEntities {
		t.Errorf("got %d, want %d", destroyedEntitiesCount, ecs.MaxEntities)
	}
}

func TestCreateEntity(t *testing.T) {
	cases := []struct {
		name string
		in   int      // Number of entities to create
		want []uint32 // IDs of the created entities
	}{
		{"should start with no entities", 0, nil},
		{"should create one entity with id", 1, []uint32{0}},
		{"should create many entities with ids", 3, []uint32{0, 1, 2}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			entityManager := ecs.NewEntityManager(ecs.MaxEntities)

			entities := []ecs.Entity{}
			for i := 0; i < c.in; i++ {
				entity, _ := entityManager.CreateEntity()
				entities = append(entities, entity)
			}
			createdEntitiesCount, _ := entityManager.EntitiesCount()

			if createdEntitiesCount != c.in {
				t.Errorf("got %d, want %d", createdEntitiesCount, c.in)
			}
			for i, entity := range entities {
				if entity != c.want[i] {
					t.Errorf("got %d, want %d", entity, c.want[i])
				}
			}
		})
	}
}

func TestDestroyEntity(t *testing.T) {
	cases := []struct {
		name string
		in   int // Number of entities to create and destroy
	}{
		{"should destroy one entity", 1},
		{"should destroy many entities", 3},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			entityManager := ecs.NewEntityManager(ecs.MaxEntities)

			entities := []ecs.Entity{}
			for i := 0; i < c.in; i++ {
				entity, _ := entityManager.CreateEntity()
				entities = append(entities, entity)
			}
			createdEntitiesCount, _ := entityManager.EntitiesCount()

			if createdEntitiesCount != c.in {
				t.Errorf("got %d, want %d", createdEntitiesCount, c.in)
			}

			for i := createdEntitiesCount - 1; i >= 0; i-- {
				entityManager.DestroyEntity(entities[i])
			}
			createdEntitiesCount, _ = entityManager.EntitiesCount()

			if createdEntitiesCount != 0 {
				t.Errorf("got %d, want 0", createdEntitiesCount)
			}
		})
	}
}
