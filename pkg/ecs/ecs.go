package ecs

type ecs struct {
	componentManager *componentManager
	entityManager    *entityManager
	systemManager    *systemManager
}

func NewECS(maxEntities int) *ecs {
	return &ecs{
		componentManager: NewComponentManager(maxEntities),
		entityManager:    NewEntityManager(maxEntities),
		systemManager:    NewSystemManager(),
	}
}

func (e *ecs) AddComponent(entity Entity, component Component) error {
	err := e.componentManager.AddComponent(entity, component)

	if err == nil {
		signature := e.componentManager.Signature(entity)
		e.systemManager.UpdateEntities(entity, signature)
	}

	return err
}

func (e *ecs) CreateEntity() (Entity, error) {
	return e.entityManager.CreateEntity()
}

func (e *ecs) DestroyEntity(entity Entity) {
	e.systemManager.RemoveEntity(entity)
	e.componentManager.DestroyComponents(entity)
	e.entityManager.DestroyEntity(entity)
}

func (e *ecs) GetComponent(entity Entity, componentType ComponentType) Component {
	return e.componentManager.GetComponent(entity, componentType)
}

func (e *ecs) RegisterSystem(system System) {
	e.systemManager.RegisterSystem(system)
}

func (e *ecs) RemoveComponent(entity Entity, componentType ComponentType) {
	signature := e.componentManager.Signature(entity)
	e.systemManager.UpdateEntities(entity, signature)

	e.componentManager.RemoveComponent(entity, componentType)
}

func (e *ecs) UpdateSystems(dt float64) {
	e.systemManager.UpdateSystems(dt)
}
