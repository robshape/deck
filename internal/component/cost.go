package component

type CostComponent struct {
	Resources uint
}

func (cc *CostComponent) Type() uint64 {
	return costComponentType
}
