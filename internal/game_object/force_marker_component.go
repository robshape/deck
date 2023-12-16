package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type ForceMarkerComponent struct{}

func (fmc *ForceMarkerComponent) Type() ecs.ComponentType {
	return forceMarkerComponentType
}
