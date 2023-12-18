package game_component

import (
	"github.com/robshape/deck/internal/ecs"
)

type AttackComponent struct {
	Damage uint
}

func (ac *AttackComponent) Type() ecs.ComponentType {
	return attackComponentType
}
