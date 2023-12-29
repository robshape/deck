package render

import (
	"github.com/robshape/deck/internal/global"
)

type RenderComponent struct{}

func (rc *RenderComponent) Type() uint64 {
	return global.RenderComponentType
}
