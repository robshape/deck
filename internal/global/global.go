package global

const (
	// Automatically increment by the power of two
	AbilityComponentType = 1 << iota
	AttackComponentType
	CardComponentType
	CostComponentType
	FactionComponentType
	ForceComponentType
	HealthComponentType
	ResourcesComponentType
	TraitComponentType
)
