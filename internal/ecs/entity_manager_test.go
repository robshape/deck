package ecs

import (
	"testing"
)

func TestNewEntityManager(t *testing.T) {
	entityManager := NewEntityManger()

	if entityManager == nil {
		t.Error("should be non-nil")
	}
}

func TestCreateEntity(t *testing.T) {
	cases := []struct {
		name             string
		numberOfEntities int
		entityIds        []uint32
	}{
		{"should start with no entities", 0, nil},
		{"should create one entity with id", 1, []uint32{1}},
		{"should create many entities with ids", 3, []uint32{1, 2, 3}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			entityManager := NewEntityManger()

			for i := 0; i < c.numberOfEntities; i++ {
				entityManager.CreateEntity()
			}
			entities := entityManager.Entities()

			if len(entities) != c.numberOfEntities {
				t.Error(c.name)
			}
			for i, e := range entities {
				if e.id != c.entityIds[i] {
					t.Error(c.name)
				}
			}
		})
	}
}

func TestDestroyEntity(t *testing.T) {
	cases := []struct {
		name             string
		numberOfEntities int
	}{
		{"should destroy one entity", 1},
		{"should destroy many entities", 3},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			entityManager := NewEntityManger()

			for i := 0; i < c.numberOfEntities; i++ {
				entityManager.CreateEntity()
			}
			entities := entityManager.Entities()

			if len(entities) != c.numberOfEntities {
				t.Error(c.name)
			}

			for i := len(entities) - 1; i >= 0; i-- {
				entityManager.DestroyEntity(entities[i])
			}
			entities = entityManager.Entities()

			if len(entities) != 0 {
				t.Error(c.name)
			}
		})
	}
}
