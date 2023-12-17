package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type propertiesComponent struct {
	ability   func()
	attack    uint
	cost      uint
	force     uint
	resources uint
}

func (pc *propertiesComponent) Type() ecs.ComponentType {
	return propertiesComponentType
}
