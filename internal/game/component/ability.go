package component

import (
	"github.com/robshape/deck/internal/global"
)

type AbilityComponent struct {
	Ability func()
}

func (ac *AbilityComponent) Type() uint64 {
	return global.AbilityComponentType
}
