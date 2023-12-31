package component

import (
	"github.com/robshape/deck/internal/global"
)

type trait int

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

func (tc *TraitComponent) Type() uint64 {
	return global.TraitComponentType
}
