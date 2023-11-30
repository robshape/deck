package ecs

import (
	"sync/atomic"
)

type entityManager struct {
	counter  uint32
	entities []*entity
}

func NewEntityManger() *entityManager {
	return &entityManager{
		entities: []*entity{},
	}
}

func (p *entityManager) CreateEntity() *entity {
	entity := &entity{
		id: atomic.AddUint32(&p.counter, 1),
	}
	p.entities = append(p.entities, entity)
	return entity
}

func (p *entityManager) DestroyEntity(entity *entity) {
	for i, e := range p.entities {
		if e.id == entity.id {
			p.entities = append(p.entities[:i], p.entities[i+1:]...)
			return
		}
	}
}
