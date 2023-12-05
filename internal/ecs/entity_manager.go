package ecs

type entityManager struct {
	createdEntitiesCount int
	destroyedEntities    []entity
}

func newEntityManager(size int) *entityManager {
	preallocated := make([]entity, size)
	for i := uint32(0); i < uint32(size); i++ {
		preallocated[i] = i
	}

	return &entityManager{
		destroyedEntities: preallocated,
	}
}

func (em *entityManager) createEntity() entity {
	entity := em.destroyedEntities[0]
	em.destroyedEntities = em.destroyedEntities[1:]

	em.createdEntitiesCount++

	return entity
}

func (em *entityManager) destroyEntity(entity entity) {
	em.destroyedEntities = append(em.destroyedEntities, entity)

	em.createdEntitiesCount--
}

func (em *entityManager) entitiesCount() (createdEntitiesCount int, destroyedEntitiesCount int) {
	return em.createdEntitiesCount, len(em.destroyedEntities)
}
