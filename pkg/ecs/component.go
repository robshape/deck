package ecs

const maxComponents = 64 // Limited by uint64 bitmask

type Component interface {
	Type() ComponentType
}

type ComponentType = uint64

type Signature = ComponentType // Bitmask of all components added to an entity
