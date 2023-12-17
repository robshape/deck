package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type reward struct {
	ability   func()
	force     uint
	resources uint
}

type healthComponent struct {
	hitPoints uint
	reward    reward
}

func (hc *healthComponent) Type() ecs.ComponentType {
	return healthComponentType
}
