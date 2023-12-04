package ecs

type entityManager struct {
	activeEntitiesCount int
	inactiveEntities    []entity
}

func NewEntityManager(size int) *entityManager {
	allocatedSlice := make([]entity, size)
	for i := uint32(0); i < uint32(size); i++ {
		allocatedSlice[i] = i
	}

	return &entityManager{
		inactiveEntities: allocatedSlice,
	}
}

func (em *entityManager) CreateEntity() entity {
	entity := em.inactiveEntities[0]
	em.inactiveEntities = em.inactiveEntities[1:]

	em.activeEntitiesCount++

	return entity
}

func (em *entityManager) DestroyEntity(entity entity) {
	em.inactiveEntities = append(em.inactiveEntities, entity)

	em.activeEntitiesCount--
}

func (em *entityManager) EntitiesCount() (int, int) {
	return em.activeEntitiesCount, len(em.inactiveEntities)
}
