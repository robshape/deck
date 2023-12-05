package ecs

const maxComponents = 64 // Limited by uint64 bitmask

type ComponentType uint64

type ComponentsMask = ComponentType

type component interface {
	Type() ComponentType
}
