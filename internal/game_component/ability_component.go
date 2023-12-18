package game_component

import (
	"github.com/robshape/deck/internal/ecs"
)

type AbilityComponent struct {
	Ability func()
}

func (ac *AbilityComponent) Type() ecs.ComponentType {
	return abilityComponentType
}
