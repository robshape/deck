package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type cardType uint

const (
	cardTypeBase cardType = iota
	cardTypeCapitalShip
	cardTypeUnit
)

type factionSymbol uint

const (
	factionSymbolEmpire factionSymbol = iota
	factionSymbolNeutral
	factionSymbolRebel
)

type trait uint

const (
	traitsBountyHunter trait = iota
	traitsDroid
	traitsFighter
	traitsOfficer
	traitsTrooper
)

type descriptionComponent struct {
	cardType      cardType
	factionSymbol factionSymbol
	title         string
	traits        []trait
	uniqueSymbol  bool
}

func (dc *descriptionComponent) Type() ecs.ComponentType {
	return descriptionComponentType
}
