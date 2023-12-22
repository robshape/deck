package component

type faction int

const (
	FactionEmpire faction = iota
	FactionNeutral
	FactionRebel
)

type FactionComponent struct {
	Faction faction
}

func (fc *FactionComponent) Type() uint64 {
	return factionComponentType
}
