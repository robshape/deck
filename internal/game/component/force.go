package game_component

import (
	"github.com/robshape/deck/pkg/ecs"
)

type ForceComponent struct {
	Force uint
}

func (fc *ForceComponent) Type() ecs.ComponentType {
	return forceComponentType
}
