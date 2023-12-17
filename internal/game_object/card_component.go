package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type CardType uint

const (
	CardTypeBase CardType = iota
	CardTypeCapitalShip
	CardTypeUnit
)

type FactionSymbol uint

const (
	FactionSymbolEmpire FactionSymbol = iota
	FactionSymbolNeutral
	FactionSymbolRebel
)

type Trait uint

const (
	TraitsBountyHunter Trait = iota
	TraitsDroid
	TraitsFighter
	TraitsOfficer
	TraitsTrooper
)

type CardComponent struct {
	ability       func()
	cardType      CardType
	factionSymbol FactionSymbol
	title         string
	traits        []Trait
	uniqueSymbol  bool
}

func (cc *CardComponent) Type() ecs.ComponentType {
	return cardComponentType
}
