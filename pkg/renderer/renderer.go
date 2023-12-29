package renderer

import (
	"fmt"
)

type Renderer interface {
	Render(format string, a ...any)
}

type renderer struct{}

func NewRenderer() Renderer {
	return &renderer{}
}

func (r *renderer) Render(format string, a ...any) {
	fmt.Printf(format, a...)
}
