package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type MarkerType uint

const (
	MarkerTypeDamageCounter MarkerType = iota
	MarkerTypeForceMarker
	MarkerTypeResourceCounter
)

type MarkerComponent struct {
	markerType MarkerType
}

func (mc *MarkerComponent) Type() ecs.ComponentType {
	return markerComponentType
}
