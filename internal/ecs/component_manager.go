package ecs

import (
	"errors"
)

type componentManager struct {
	entityComponents []*entityComponents
}

type entityComponents struct {
	components     []component
	componentsMask ComponentsMask // Bitmask of all added components
}

func NewComponentManager(maxEntities int) *componentManager {
	preallocated := make([]*entityComponents, maxEntities)
	for i := range preallocated {
		preallocated[i] = &entityComponents{}
	}

	return &componentManager{
		entityComponents: preallocated,
	}
}

func (cm *componentManager) AddComponent(entity Entity, component component) error {
	ec := cm.entityComponents[entity]

	if len(ec.components) == maxComponents {
		return errors.New("maximum number of components already added")
	} else if (ec.componentsMask & component.Type()) == component.Type() {
		return errors.New("component already added")
	}

	ec.components = append(ec.components, component)
	ec.componentsMask |= component.Type() // Add component to bitmask
	return nil
}

func (cm *componentManager) ComponentsCount(entity Entity) int {
	return len(cm.entityComponents[entity].components)
}

func (cm *componentManager) ComponentsMask(entity Entity) ComponentsMask {
	return cm.entityComponents[entity].componentsMask
}

func (cm *componentManager) DestroyEntityComponents(entity Entity) {
	cm.entityComponents[entity] = &entityComponents{}
}

func (cm *componentManager) EntityComponentsCount() int {
	return len(cm.entityComponents)
}

func (cm *componentManager) GetComponent(entity Entity, componentType ComponentType) component {
	ec := cm.entityComponents[entity]

	for _, component := range ec.components {
		if component.Type() == componentType {
			return component
		}
	}

	return nil
}

func (cm *componentManager) RemoveComponent(entity Entity, componentType ComponentType) {
	ec := cm.entityComponents[entity]

	for i, component := range ec.components {
		if component.Type() == componentType {
			ec.components[i] = ec.components[len(ec.components)-1]
			ec.components[len(ec.components)-1] = nil
			ec.components = ec.components[:len(ec.components)-1]

			ec.componentsMask ^= componentType // Remove component from bitmask

			break
		}
	}
}
