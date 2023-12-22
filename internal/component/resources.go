package component

type ResourcesComponent struct {
	Resources uint
}

func (rc *ResourcesComponent) Type() uint64 {
	return resourcesComponentType
}
