package ecs

import (
	"sync/atomic"
)

type entityManager struct {
	counter  uint32
	entities []*entity
}

func NewEntityManager() *entityManager {
	return &entityManager{
		entities: []*entity{},
	}
}

func (em *entityManager) CreateEntity() *entity {
	entity := &entity{
		id: atomic.AddUint32(&em.counter, 1),
	}
	em.entities = append(em.entities, entity)
	return entity
}

func (em *entityManager) DestroyEntity(entity *entity) {
	for i, e := range em.entities {
		if e.id == entity.id {
			em.entities[i] = em.entities[len(em.entities)-1]
			em.entities[len(em.entities)-1] = nil
			em.entities = em.entities[:len(em.entities)-1]
			break
		}
	}
}

func (em *entityManager) Entities() []*entity {
	return em.entities
}
