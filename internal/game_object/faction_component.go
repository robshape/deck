package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type faction uint

const (
	factionEmpire faction = iota
	factionNeutral
	factionRebel
)

type factionComponent struct {
	faction faction
}

func (fc *factionComponent) Type() ecs.ComponentType {
	return factionComponentType
}
