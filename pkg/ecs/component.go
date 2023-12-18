package ecs

const maxComponents = 64 // Limited by uint64 bitmask

type ComponentType uint64

type Component interface {
	Type() ComponentType
}

type Signature = ComponentType // Bitmask of all components added to an entity
