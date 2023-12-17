package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type abilityComponent struct {
	ability func()
}

func (ac *abilityComponent) Type() ecs.ComponentType {
	return abilityComponentType
}
