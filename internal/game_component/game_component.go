package game_component

import (
	"github.com/robshape/deck/internal/ecs"
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
