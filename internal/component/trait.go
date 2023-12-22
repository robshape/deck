package component

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

func (tc *TraitComponent) Type() uint64 {
	return traitComponentType
}
