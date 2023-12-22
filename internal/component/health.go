package component

type Reward struct {
	Ability   func()
	Force     int
	Resources int
}

type HealthComponent struct {
	HitPoints int
	Reward    Reward
}

func (hc *HealthComponent) Type() uint64 {
	return healthComponentType
}
