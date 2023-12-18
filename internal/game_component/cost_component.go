package game_component

import (
	"github.com/robshape/deck/internal/ecs"
)

type CostComponent struct {
	Resources uint
}

func (cc *CostComponent) Type() ecs.ComponentType {
	return costComponentType
}
