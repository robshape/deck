package ecs

type System interface {
	AddEntity(entity Entity)
	Entities() []Entity
	RemoveEntity(entity Entity)
	Signature() Signature
	Update(dt float64)
}
