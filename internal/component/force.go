package component

type ForceComponent struct {
	Force int
}

func (fc *ForceComponent) Type() uint64 {
	return forceComponentType
}
