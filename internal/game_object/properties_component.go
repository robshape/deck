package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type PropertiesComponent struct {
	attack    uint
	cost      uint
	force     uint
	resources uint
}

func (pc *PropertiesComponent) Type() ecs.ComponentType {
	return propertiesComponentType
}
