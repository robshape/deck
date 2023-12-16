package game_object

import (
	"github.com/robshape/deck/internal/ecs"
)

type DamageCounterComponent struct{}

func (dcc *DamageCounterComponent) Type() ecs.ComponentType {
	return damageCounterComponentType
}
