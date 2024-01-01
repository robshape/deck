package game

import (
	"github.com/robshape/deck/pkg/ecs"
)

type ecsManager interface {
	AddComponent(entity uint32, component ecs.Component) error
	CreateEntity() (uint32, error)
	DestroyEntity(entity uint32)
	RegisterSystem(system ecs.System)
	UpdateSystems(dt float64)
}
