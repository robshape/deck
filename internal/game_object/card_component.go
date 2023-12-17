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

type designation uint

const (
	designationEmpireStarter designation = iota
	designationGalaxyDeck
	designationOuterRimPilotDeck
	designationRebelStarter
)

type cardComponent struct {
	cardType    cardType
	description string
	designation designation
	title       string
	unique      bool
}

func (cc *cardComponent) Type() ecs.ComponentType {
	return cardComponentType
}
