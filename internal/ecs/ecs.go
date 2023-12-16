package ecs

type ecsManager struct {
	componentManager *componentManager
	entityManager    *entityManager
	systemManager    *systemManager
}

func NewECSManager(maxEntities int) *ecsManager {
	return &ecsManager{
		componentManager: NewComponentManager(maxEntities),
		entityManager:    NewEntityManager(maxEntities),
		systemManager:    NewSystemManager(),
	}
}

func (em *ecsManager) AddComponent(entity Entity, component component) error {
	return em.componentManager.AddComponent(entity, component)
}

func (em *ecsManager) CreateEntity() (Entity, error) {
	return em.entityManager.CreateEntity()
}

func (em *ecsManager) DestroyEntity(entity Entity) {
	em.componentManager.DestroyEntityComponents(entity)
	em.entityManager.DestroyEntity(entity)
}

func (em *ecsManager) GetComponent(entity Entity, componentType ComponentType) component {
	return em.componentManager.GetComponent(entity, componentType)
}

func (em *ecsManager) RegisterSystem(system System) {
	em.systemManager.RegisterSystem(system)
}

func (em *ecsManager) RemoveComponent(entity Entity, componentType ComponentType) {
	em.componentManager.RemoveComponent(entity, componentType)
}

func (em *ecsManager) UpdateSystems(dt float64) {
	em.systemManager.UpdateSystems(dt)
}
