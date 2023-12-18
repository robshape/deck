package game_component

import (
	"github.com/robshape/deck/internal/ecs"
)

type ResourcesComponent struct {
	Resources uint
}

func (rc *ResourcesComponent) Type() ecs.ComponentType {
	return resourcesComponentType
}
