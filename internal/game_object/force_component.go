package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type forceComponent struct {
	force uint
}

func (fc *forceComponent) Type() ecs.ComponentType {
	return forceComponentType
}
