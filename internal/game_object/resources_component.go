package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type resourcesComponent struct {
	resources uint
}

func (rc *resourcesComponent) Type() ecs.ComponentType {
	return resourcesComponentType
}
