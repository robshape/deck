package ecs

type EcsManager struct {
	componentManager *componentManager
	entityManager    *entityManager
	systemManager    *systemManager
}

func NewECSManager(maxEntities int) *EcsManager {
	return &EcsManager{
		componentManager: NewComponentManager(maxEntities),
		entityManager:    NewEntityManager(maxEntities),
		systemManager:    NewSystemManager(),
	}
}

func (em *EcsManager) AddComponent(entity Entity, component Component) error {
	return em.componentManager.AddComponent(entity, component)
}

func (em *EcsManager) CreateEntity() (Entity, error) {
	return em.entityManager.CreateEntity()
}

func (em *EcsManager) DestroyEntity(entity Entity) {
	em.componentManager.DestroyEntityComponents(entity)
	em.entityManager.DestroyEntity(entity)
}

func (em *EcsManager) GetComponent(entity Entity, componentType ComponentType) Component {
	return em.componentManager.GetComponent(entity, componentType)
}

func (em *EcsManager) RegisterSystem(system System) {
	em.systemManager.RegisterSystem(system)
}

func (em *EcsManager) RemoveComponent(entity Entity, componentType ComponentType) {
	em.componentManager.RemoveComponent(entity, componentType)
}

func (em *EcsManager) UpdateSystems(dt float64) {
	em.systemManager.UpdateSystems(dt)
}
