package component

type ResourcesComponent struct {
	Resources int
}

func (rc *ResourcesComponent) Type() uint64 {
	return resourcesComponentType
}
