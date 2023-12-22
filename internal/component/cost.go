package component

type CostComponent struct {
	Resources int
}

func (cc *CostComponent) Type() uint64 {
	return costComponentType
}
