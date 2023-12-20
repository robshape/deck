package component

import (
	"github.com/robshape/deck/pkg/ecs"
)

type Reward struct {
	Ability   func()
	Force     uint
	Resources uint
}

type HealthComponent struct {
	HitPoints uint
	Reward    Reward
}

func (hc *HealthComponent) Type() ecs.ComponentType {
	return healthComponentType
}
