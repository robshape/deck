package ecs

type ComponentMask uint64

type Component interface {
	Mask() ComponentMask
}
