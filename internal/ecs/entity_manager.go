package ecs

type entityManager struct {
	activeEntitiesCount int
	inactiveEntities    []entity
}

func newEntityManager(size int) *entityManager {
	preallocated := make([]entity, size)
	for i := uint32(0); i < uint32(size); i++ {
		preallocated[i] = i
	}

	return &entityManager{
		inactiveEntities: preallocated,
	}
}

func (em *entityManager) createEntity() entity {
	entity := em.inactiveEntities[0]
	em.inactiveEntities = em.inactiveEntities[1:]

	em.activeEntitiesCount++

	return entity
}

func (em *entityManager) destroyEntity(entity entity) {
	em.inactiveEntities = append(em.inactiveEntities, entity)

	em.activeEntitiesCount--
}

func (em *entityManager) entitiesCount() (int, int) {
	return em.activeEntitiesCount, len(em.inactiveEntities)
}
