package game_component

import (
	"github.com/robshape/deck/pkg/ecs"
)

type faction uint

const (
	FactionEmpire faction = iota
	FactionNeutral
	FactionRebel
)

type FactionComponent struct {
	Faction faction
}

func (fc *FactionComponent) Type() ecs.ComponentType {
	return factionComponentType
}
