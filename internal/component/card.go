package component

type cardType uint

const (
	CardTypeBase cardType = iota
	CardTypeCapitalShip
	CardTypeUnit
)

type designation uint

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
	return cardComponentType
}
