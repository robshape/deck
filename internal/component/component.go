package component

import (
	"github.com/robshape/deck/pkg/ecs"
)

const (
	// Automatically increment by the power of two
	abilityComponentType ecs.ComponentType = 1 << iota
	attackComponentType
	cardComponentType
	costComponentType
	factionComponentType
	forceComponentType
	healthComponentType
	resourcesComponentType
	traitComponentType
)
