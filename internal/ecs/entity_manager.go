package ecs

type entityManager struct {
	createdEntitiesCount int
	destroyedEntities    []Entity
}

func NewEntityManager(size int) *entityManager {
	preallocated := make([]Entity, size)
	for i := uint32(0); i < uint32(size); i++ {
		preallocated[i] = i
	}

	return &entityManager{
		destroyedEntities: preallocated,
	}
}

func (em *entityManager) CreateEntity() Entity {
	entity := em.destroyedEntities[0]
	em.destroyedEntities = em.destroyedEntities[1:]

	em.createdEntitiesCount++

	return entity
}

func (em *entityManager) DestroyEntity(entity Entity) {
	em.destroyedEntities = append(em.destroyedEntities, entity)

	em.createdEntitiesCount--
}

func (em *entityManager) EntitiesCount() (createdEntitiesCount int, destroyedEntitiesCount int) {
	return em.createdEntitiesCount, len(em.destroyedEntities)
}
