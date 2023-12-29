package ecs

type System interface {
	AddEntity(entity Entity)
	RemoveEntity(entity Entity)
	Signature() Signature
	Update(dt float64)
}
