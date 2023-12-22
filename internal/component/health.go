package component

type Reward struct {
	Ability   func()
	Force     uint
	Resources uint
}

type HealthComponent struct {
	HitPoints uint
	Reward    Reward
}

func (hc *HealthComponent) Type() uint64 {
	return healthComponentType
}
