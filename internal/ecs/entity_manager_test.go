package ecs

import (
	"testing"
)

func TestNewEntityManager(t *testing.T) {
	entityManager := NewEntityManager()

	if entityManager == nil {
		t.Error("got nil, want non-nil")
	}
}

func TestCreateEntity(t *testing.T) {
	cases := []struct {
		name string
		in   int      // Number of entities to create
		want []uint32 // IDs of the created entities
	}{
		{"should start with no entities", 0, nil},
		{"should create one entity with id", 1, []uint32{1}},
		{"should create many entities with ids", 3, []uint32{1, 2, 3}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			entityManager := NewEntityManager()

			for i := 0; i < c.in; i++ {
				entityManager.CreateEntity()
			}
			entities := entityManager.Entities()

			if len(entities) != c.in {
				t.Errorf("got %d, want %d", len(entities), c.in)
			}
			for i, e := range entities {
				if e.id != c.want[i] {
					t.Errorf("got %d, want %d", e.id, c.want[i])
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
			entityManager := NewEntityManager()

			for i := 0; i < c.in; i++ {
				entityManager.CreateEntity()
			}
			entities := entityManager.Entities()

			if len(entities) != c.in {
				t.Errorf("got %d, want %d", len(entities), c.in)
			}

			for i := len(entities) - 1; i >= 0; i-- {
				entityManager.DestroyEntity(entities[i])
			}
			entities = entityManager.Entities()

			if len(entities) != 0 {
				t.Errorf("got %d, want 0", len(entities))
			}
		})
	}
}
