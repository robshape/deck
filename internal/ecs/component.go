package ecs

const maxComponents = 64 // Limited by uint64 bitmask

type ComponentType uint64

type component interface {
	Type() ComponentType
}

type componentsMask = ComponentType
