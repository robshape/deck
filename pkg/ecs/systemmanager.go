package ecs

type systemManager struct {
	systems []System
}

func NewSystemManager() *systemManager {
	return &systemManager{
		systems: []System{},
	}
}

func (sm *systemManager) RegisterSystem(system System) {
	sm.systems = append(sm.systems, system)
}

func (sm *systemManager) RemoveEntity(entity Entity) {
	for _, system := range sm.systems {
		system.RemoveEntity(entity)
	}
}

func (sm *systemManager) SystemsCount() int {
	return len(sm.systems)
}

func (sm *systemManager) UpdateEntities(entity Entity, signature Signature) {
	for _, system := range sm.systems {
		if signature&system.Signature() == system.Signature() {
			system.AddEntity(entity)
		} else {
			system.RemoveEntity(entity)
		}
	}
}

func (sm *systemManager) UpdateSystems(dt float64) {
	for _, system := range sm.systems {
		system.Update(dt)
	}
}
