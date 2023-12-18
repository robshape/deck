package game_component

import (
	"github.com/robshape/deck/internal/ecs"
)

type Reward struct {
	Ability   func()
	Force     uint
	Resources uint
}

type healthComponent struct {
	HitPoints uint
	Reward    Reward
}

func (hc *healthComponent) Type() ecs.ComponentType {
	return healthComponentType
}
