package ecs

import (
	"errors"
)

type componentManager struct {
	entityComponents []*entityComponents
}

type entityComponents struct {
	components     []Component
	componentsMask ComponentsMask // Bitmask of all added components
}

func NewComponentManager(size int) *componentManager {
	allocatedSlice := make([]*entityComponents, size)
	for i := range allocatedSlice {
		allocatedSlice[i] = &entityComponents{}
	}

	return &componentManager{
		entityComponents: allocatedSlice,
	}
}

func (cm *componentManager) AddComponent(entity entity, component Component) error {
	ec := cm.entityComponents[entity]

	if len(ec.components) == MAX_COMPONENTS {
		return errors.New("maximum number of components already added")
	} else if (ec.componentsMask & component.Type()) == component.Type() {
		return errors.New("component already added")
	}

	ec.components = append(ec.components, component)
	ec.componentsMask |= component.Type() // Add component to bitmask
	return nil
}

func (cm *componentManager) ComponentsCount(entity entity) int {
	ec := cm.entityComponents[entity]
	return len(ec.components)
}

func (cm *componentManager) ComponentsMask(entity entity) ComponentsMask {
	ec := cm.entityComponents[entity]
	return ec.componentsMask
}

func (cm *componentManager) RemoveComponent(entity entity, componentType ComponentType) {
	ec := cm.entityComponents[entity]

	for i, c := range ec.components {
		if c.Type() == componentType {
			ec.components[i] = ec.components[len(ec.components)-1]
			ec.components[len(ec.components)-1] = nil
			ec.components = ec.components[:len(ec.components)-1]

			ec.componentsMask ^= componentType // Remove component from bitmask

			break
		}
	}
}
