package ecs

type ecs struct {
	componentManager *componentManager
	entityManager    *entityManager
}

func NewECS(size int) *ecs {
	return &ecs{
		componentManager: newComponentManager(size),
		entityManager:    newEntityManager(size),
	}
}

func (ecs *ecs) AddComponent(entity entity, component component) error {
	return ecs.componentManager.addComponent(entity, component)
}

func (ecs *ecs) CreateEntity() entity {
	return ecs.entityManager.createEntity()
}

func (ecs *ecs) DestroyEntity(entity entity) {
	ecs.componentManager.destroyEntityComponents(entity)
	ecs.entityManager.destroyEntity(entity)
}

func (ecs *ecs) RemoveComponent(entity entity, componentType ComponentType) {
	ecs.componentManager.removeComponent(entity, componentType)
}
