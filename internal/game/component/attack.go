package component

import (
	"github.com/robshape/deck/internal/global"
)

type AttackComponent struct {
	Damage int
}

func (ac *AttackComponent) Type() uint64 {
	return global.AttackComponentType
}
