package component

import (
	"github.com/robshape/deck/pkg/ecs"
)

type trait uint

const (
	TraitBountyHunter trait = iota
	TraitDroid
	TraitFighter
	TraitOfficer
	TraitTrooper
)

type TraitComponent struct {
	Traits []trait
}

func (tc *TraitComponent) Type() ecs.ComponentType {
	return traitComponentType
}
