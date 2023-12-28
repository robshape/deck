package component

import (
	"github.com/robshape/deck/internal/global"
)

type CostComponent struct {
	Resources int
}

func (cc *CostComponent) Type() uint64 {
	return global.CostComponentType
}
