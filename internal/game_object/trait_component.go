package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type trait uint

const (
	traitBountyHunter trait = iota
	traitDroid
	traitFighter
	traitOfficer
	traitTrooper
)

type traitComponent struct {
	traits []trait
}

func (tc *traitComponent) Type() ecs.ComponentType {
	return traitComponentType
}
