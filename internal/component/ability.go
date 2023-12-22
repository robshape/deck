package component

type AbilityComponent struct {
	Ability func()
}

func (ac *AbilityComponent) Type() uint64 {
	return abilityComponentType
}
