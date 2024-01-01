package ecs

type System interface {
	AddEntity(entity Entity)
	RemoveEntity(entity Entity)
	Entities() []Entity
	Signature() Signature
	Update(dt float64)
}
