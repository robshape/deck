package ecs

const (
	// Automatically increment by the power of two
	COMPONENT_DAMAGE_COUNTER ComponentMask = 1 << iota
	COMPONENT_FORCE_MARKER
	COMPONENT_RESOURCE_COUNTER
)

type ComponentMask uint64

type Component interface {
	Mask() ComponentMask
}
