package ecs

import (
	"math"
	"testing"
)

type mockComponent struct {
	mask ComponentMask
}

func (mc *mockComponent) Mask() ComponentMask {
	return mc.mask
}

func TestAddComponent(t *testing.T) {
	cases := []struct {
		name string
		in   []ComponentMask
		want ComponentMask
	}{
		{"should add one component with mask", []ComponentMask{COMPONENT_DAMAGE_COUNTER}, COMPONENT_DAMAGE_COUNTER},
		{"should add many components with mask", []ComponentMask{COMPONENT_DAMAGE_COUNTER, COMPONENT_FORCE_MARKER, COMPONENT_RESOURCE_COUNTER}, COMPONENT_DAMAGE_COUNTER | COMPONENT_FORCE_MARKER | COMPONENT_RESOURCE_COUNTER},
		{"should not add duplicate component with mask", []ComponentMask{COMPONENT_DAMAGE_COUNTER, COMPONENT_DAMAGE_COUNTER, COMPONENT_DAMAGE_COUNTER}, COMPONENT_DAMAGE_COUNTER},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			entityManager := NewEntityManager()
			entity := entityManager.CreateEntity()

			for _, componentTypeMask := range c.in {
				entity.AddComponent(&mockComponent{mask: componentTypeMask})
			}
			count := math.Round(math.Sqrt(float64(c.want)))

			if len(entity.Components()) != int(count) {
				t.Errorf("got %d, want %d", len(entity.Components()), int(count))
			}
			if entity.ComponentsMask() != c.want {
				t.Errorf("got %d, want %d", entity.ComponentsMask(), c.want)
			}
		})
	}
}

func TestRemoveComponent(t *testing.T) {
	cases := []struct {
		name string
		in   []ComponentMask
	}{
		{"should remove one component", []ComponentMask{COMPONENT_DAMAGE_COUNTER}},
		{"should remove many components", []ComponentMask{COMPONENT_DAMAGE_COUNTER, COMPONENT_FORCE_MARKER, COMPONENT_RESOURCE_COUNTER}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			entityManager := NewEntityManager()
			entity := entityManager.CreateEntity()

			for _, componentTypeMask := range c.in {
				entity.AddComponent(&mockComponent{mask: componentTypeMask})
			}

			for _, componentTypeMask := range c.in {
				entity.RemoveComponent(componentTypeMask)
			}

			if len(entity.Components()) != 0 {
				t.Errorf("got %d, want 0", len(entity.Components()))
			}
			if entity.ComponentsMask() != 0 {
				t.Errorf("got %d, want 0", entity.ComponentsMask())
			}
		})
	}
}
