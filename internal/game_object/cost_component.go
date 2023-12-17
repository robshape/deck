package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type costComponent struct {
	resources uint
}

func (cc *costComponent) Type() ecs.ComponentType {
	return costComponentType
}
