package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type ResourceCounterComponent struct{}

func (rcc *ResourceCounterComponent) Type() ecs.ComponentType {
	return forceMarkerComponentType
}
