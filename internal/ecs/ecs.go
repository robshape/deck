package ecs

type ecsCoordinator struct {
	componentManager *componentManager
	entityManager    *entityManager
}

func NewECSCoordinator(maxEntities int) *ecsCoordinator {
	return &ecsCoordinator{
		componentManager: NewComponentManager(maxEntities),
		entityManager:    NewEntityManager(maxEntities),
	}
}

func (ec *ecsCoordinator) AddComponent(entity Entity, component component) error {
	return ec.componentManager.AddComponent(entity, component)
}

func (ec *ecsCoordinator) CreateEntity() (Entity, error) {
	return ec.entityManager.CreateEntity()
}

func (ec *ecsCoordinator) DestroyEntity(entity Entity) {
	ec.componentManager.DestroyEntityComponents(entity)
	ec.entityManager.DestroyEntity(entity)
}

func (ec *ecsCoordinator) RemoveComponent(entity Entity, componentType ComponentType) {
	ec.componentManager.RemoveComponent(entity, componentType)
}
