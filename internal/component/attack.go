package component

type AttackComponent struct {
	Damage int
}

func (ac *AttackComponent) Type() uint64 {
	return attackComponentType
}
