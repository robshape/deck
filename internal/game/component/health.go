package component

import (
	"github.com/robshape/deck/internal/global"
)

type Reward struct {
	Ability   func()
	Force     int
	Resources int
}

type HealthComponent struct {
	HitPoints int
	Reward    Reward
}

func (hc *HealthComponent) Type() uint64 {
	return global.HealthComponentType
}
