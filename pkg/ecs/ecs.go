package ecs

type EcsManager interface {
	AddComponent(entity Entity, component Component) error
	CreateEntity() (Entity, error)
	DestroyEntity(entity Entity)
	GetComponent(entity Entity, componentType ComponentType) Component
	RegisterSystem(system System)
	RemoveComponent(entity Entity, componentType ComponentType)
	UpdateSystems(dt float64)
}

type ecsManager struct {
	componentManager *componentManager
	entityManager    *entityManager
	systemManager    *systemManager
}

func NewEcsManager(maxEntities int) EcsManager {
	return &ecsManager{
		componentManager: NewComponentManager(maxEntities),
		entityManager:    NewEntityManager(maxEntities),
		systemManager:    NewSystemManager(),
	}
}

func (em *ecsManager) AddComponent(entity Entity, component Component) error {
	err := em.componentManager.AddComponent(entity, component)

	if err == nil {
		signature := em.componentManager.Signature(entity)
		em.systemManager.UpdateEntities(entity, signature)
	}

	return err
}

func (em *ecsManager) CreateEntity() (Entity, error) {
	return em.entityManager.CreateEntity()
}

func (em *ecsManager) DestroyEntity(entity Entity) {
	em.systemManager.RemoveEntity(entity)
	em.componentManager.DestroyComponents(entity)
	em.entityManager.DestroyEntity(entity)
}

func (em *ecsManager) GetComponent(entity Entity, componentType ComponentType) Component {
	return em.componentManager.GetComponent(entity, componentType)
}

func (em *ecsManager) RegisterSystem(system System) {
	em.systemManager.RegisterSystem(system)
}

func (em *ecsManager) RemoveComponent(entity Entity, componentType ComponentType) {
	signature := em.componentManager.Signature(entity)
	em.systemManager.UpdateEntities(entity, signature)

	em.componentManager.RemoveComponent(entity, componentType)
}

func (em *ecsManager) UpdateSystems(dt float64) {
	em.systemManager.UpdateSystems(dt)
}
