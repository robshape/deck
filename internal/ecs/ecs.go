package ecs

type ecs struct {
	componentManager *componentManager
	entityManager    *entityManager
}

func NewECS(size int) *ecs {
	return &ecs{
		componentManager: NewComponentManager(size),
		entityManager:    NewEntityManager(size),
	}
}

func (ecs *ecs) AddComponent(entity Entity, component component) error {
	return ecs.componentManager.AddComponent(entity, component)
}

func (ecs *ecs) CreateEntity() Entity {
	return ecs.entityManager.CreateEntity()
}

func (ecs *ecs) DestroyEntity(entity Entity) {
	ecs.componentManager.DestroyEntityComponents(entity)
	ecs.entityManager.DestroyEntity(entity)
}

func (ecs *ecs) RemoveComponent(entity Entity, componentType ComponentType) {
	ecs.componentManager.RemoveComponent(entity, componentType)
}
