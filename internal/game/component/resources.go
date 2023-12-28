package component

import (
	"github.com/robshape/deck/internal/global"
)

type ResourcesComponent struct {
	Resources int
}

func (rc *ResourcesComponent) Type() uint64 {
	return global.ResourcesComponentType
}
