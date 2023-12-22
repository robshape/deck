package component

type AttackComponent struct {
	Damage uint
}

func (ac *AttackComponent) Type() uint64 {
	return attackComponentType
}
