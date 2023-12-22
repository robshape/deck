package ecs

import (
	"errors"
)

type entityManager struct {
	createdEntitiesCount int
	destroyedEntities    []Entity
}

func NewEntityManager(maxEntities int) *entityManager {
	preallocated := make([]Entity, maxEntities)
	for i := uint32(0); i < uint32(maxEntities); i++ {
		preallocated[i] = i
	}

	return &entityManager{
		destroyedEntities: preallocated,
	}
}

func (em *entityManager) CreateEntity() (Entity, error) {
	if (len(em.destroyedEntities)) == 0 {
		return 0, errors.New("CreateEntity: maximum number of entities already created")
	}

	entity := em.destroyedEntities[0]
	em.destroyedEntities = em.destroyedEntities[1:]

	em.createdEntitiesCount++

	return entity, nil
}

func (em *entityManager) DestroyEntity(entity Entity) {
	em.destroyedEntities = append(em.destroyedEntities, entity)

	em.createdEntitiesCount--
}

func (em *entityManager) EntitiesCount() (createdEntitiesCount int, destroyedEntitiesCount int) {
	return em.createdEntitiesCount, len(em.destroyedEntities)
}
