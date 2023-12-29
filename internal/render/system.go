package render

import (
	"github.com/robshape/deck/internal/global"
	"github.com/robshape/deck/pkg/renderer"
)

type renderSystem struct {
	entities []uint32
	renderer renderer.Renderer
}

func NewRenderSystem(renderer renderer.Renderer) *renderSystem {
	return &renderSystem{
		renderer: renderer,
	}
}

func (rs *renderSystem) AddEntity(entity uint32) {
	rs.entities = append(rs.entities, entity)
}

func (rs *renderSystem) RemoveEntity(entity uint32) {
	for i, e := range rs.entities {
		if e == entity {
			rs.entities[i] = rs.entities[len(rs.entities)-1]
			rs.entities[len(rs.entities)-1] = 0
			rs.entities = rs.entities[:len(rs.entities)-1]
			break
		}
	}
}

func (rs *renderSystem) Signature() uint64 {
	return global.RenderComponentType
}

func (rs *renderSystem) Update(dt float64) {
	if len(rs.entities) == 0 {
		return
	}

	rs.renderer.Render("RENDER ENTITIES: %d (dt: %f)\n", len(rs.entities), dt)
}
