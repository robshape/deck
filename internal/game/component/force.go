package component

import (
	"github.com/robshape/deck/internal/global"
)

type ForceComponent struct {
	Force int
}

func (fc *ForceComponent) Type() uint64 {
	return global.ForceComponentType
}
