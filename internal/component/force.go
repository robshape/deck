package component

type ForceComponent struct {
	Force uint
}

func (fc *ForceComponent) Type() uint64 {
	return forceComponentType
}
