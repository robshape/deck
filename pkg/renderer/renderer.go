package renderer

import (
	"fmt"
)

type renderer struct{}

func NewRenderer() *renderer {
	return &renderer{}
}

func (r *renderer) Render(format string, a ...any) {
	fmt.Printf(format, a...)
}
