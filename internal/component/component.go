package component

const (
	// Automatically increment by the power of two
	abilityComponentType = 1 << iota
	attackComponentType
	cardComponentType
	costComponentType
	factionComponentType
	forceComponentType
	healthComponentType
	resourcesComponentType
	traitComponentType
)
