package component

import (
	"github.com/robshape/deck/internal/global"
)

type cardType int

const (
	CardTypeBase cardType = iota
	CardTypeCapitalShip
	CardTypeUnit
)

type designation int

const (
	DesignationEmpireStarter designation = iota
	DesignationGalaxyDeck
	DesignationOuterRimPilotDeck
	DesignationRebelStarter
)

type CardComponent struct {
	CardType    cardType
	Description string
	Designation designation
	Title       string
	Unique      bool
}

func (cc *CardComponent) Type() uint64 {
	return global.CardComponentType
}
