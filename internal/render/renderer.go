package render

type renderer interface {
	Render(format string, a ...any)
}
