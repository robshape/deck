package renderer

import (
	"fmt"
)

type Renderer interface {
	Render(a ...any)
}

type renderer struct{}

func NewRenderer() Renderer {
	return &renderer{}
}

func (r *renderer) Render(a ...any) {
	fmt.Println(a...)
}
