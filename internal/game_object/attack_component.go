package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type attackComponent struct {
	damage uint
}

func (ac *attackComponent) Type() ecs.ComponentType {
	return attackComponentType
}
