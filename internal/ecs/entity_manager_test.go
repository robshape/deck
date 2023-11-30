package ecs

import (
	"testing"
)

func TestNewEntityManager(t *testing.T) {
	entityManager := NewEntityManger()

	if entityManager == nil {
		t.Error("Expected entityManager to be non-nil")
	}
	if len(entityManager.entities) != 0 {
		t.Error("Expected entityManager.entities to be empty")
	}
}

func TestCreateEntity(t *testing.T) {
	entityManager := NewEntityManger()
	entity1 := entityManager.CreateEntity()

	if len(entityManager.entities) != 1 {
		t.Error("Expected entityManager.entities to have length of 1")
	}
	if entity1.id != 1 {
		t.Error("Expected entity1.id to be 1")
	}

	entity2 := entityManager.CreateEntity()

	if len(entityManager.entities) != 2 {
		t.Error("Expected entityManager.entities to have length of 2")
	}
	if entity1.id != 1 {
		t.Error("Expected entity1.id to be 1")
	}
	if entity2.id != 2 {
		t.Error("Expected entity2.id to be 2")
	}
}

func TestDestroyEntity(t *testing.T) {
	entityManager := NewEntityManger()
	entity1 := entityManager.CreateEntity()
	entity2 := entityManager.CreateEntity()
	entity3 := entityManager.CreateEntity()

	if len(entityManager.entities) != 3 {
		t.Error("Expected entityManager.entities to have length of 3")
	}

	entityManager.DestroyEntity(entity2)

	if len(entityManager.entities) != 2 {
		t.Error("Expected entityManager.entities to have length of 2")
	}

	entityManager.DestroyEntity(entity1)
	entityManager.DestroyEntity(entity3)

	if len(entityManager.entities) != 0 {
		t.Error("Expected entityManager.entities to have length of 0")
	}
}
