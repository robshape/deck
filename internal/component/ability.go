package component

import (
	"github.com/robshape/deck/pkg/ecs"
)

type AbilityComponent struct {
	Ability func()
}

func (ac *AbilityComponent) Type() ecs.ComponentType {
	return abilityComponentType
}
