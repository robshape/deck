package ecs

import (
	"errors"
)

type componentManager struct {
	entityComponents []*entityComponents
}

type entityComponents struct {
	components     []component
	componentsMask componentsMask // Bitmask of all added components
}

func newComponentManager(size int) *componentManager {
	preallocated := make([]*entityComponents, size)
	for i := range preallocated {
		preallocated[i] = &entityComponents{}
	}

	return &componentManager{
		entityComponents: preallocated,
	}
}

func (cm *componentManager) addComponent(entity entity, component component) error {
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

func (cm *componentManager) componentsCount(entity entity) int {
	return len(cm.entityComponents[entity].components)
}

func (cm *componentManager) componentsMask(entity entity) componentsMask {
	return cm.entityComponents[entity].componentsMask
}

func (cm *componentManager) destroyEntityComponents(entity entity) {
	cm.entityComponents[entity] = &entityComponents{}
}

func (cm *componentManager) removeComponent(entity entity, componentType ComponentType) {
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
