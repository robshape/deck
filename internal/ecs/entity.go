package ecs

type entity struct {
	components     []Component
	componentsMask ComponentMask // Bitmask of all added components
	id             uint32
}

func (e *entity) AddComponent(component Component) {
	if (e.componentsMask & component.Mask()) == component.Mask() {
		return // Component already added
	}

	e.components = append(e.components, component)

	e.componentsMask |= component.Mask() // Add component to bitmask
}

func (e *entity) Components() []Component {
	return e.components
}

func (e *entity) ComponentsMask() ComponentMask {
	return e.componentsMask
}

func (e *entity) RemoveComponent(componentMask ComponentMask) {
	for i, c := range e.components {
		if c.Mask() == componentMask {
			e.components[i] = e.components[len(e.components)-1]
			e.components[len(e.components)-1] = nil
			e.components = e.components[:len(e.components)-1]

			e.componentsMask ^= componentMask // Remove component from bitmask

			break
		}
	}
}
