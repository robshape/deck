package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type healthComponent struct {
	hitPoints int
	reward    func()
}

func (hc *healthComponent) Type() ecs.ComponentType {
	return healthComponentType
}
