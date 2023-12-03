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

			for _, componentMask := range c.in {
				entity.AddComponent(&mockComponent{mask: componentMask})
			}
			componentsCount := len(entity.Components())
			componentsMask := entity.ComponentsMask()
			wantCount := int(math.Round(math.Sqrt(float64(c.want)))) // Count bits set in mask

			if componentsCount != wantCount {
				t.Errorf("got %d, want %d", componentsCount, wantCount)
			}
			if componentsMask != c.want {
				t.Errorf("got %d, want %d", componentsMask, c.want)
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

			for _, componentMask := range c.in {
				entity.AddComponent(&mockComponent{mask: componentMask})
			}
			componentsCount := len(entity.Components())
			inCount := len(c.in)

			if componentsCount != inCount {
				t.Errorf("got %d, want %d", componentsCount, inCount)
			}

			for _, componentMask := range c.in {
				entity.RemoveComponent(componentMask)
			}
			componentsCount = len(entity.Components())
			componentsMask := entity.ComponentsMask()

			if componentsCount != 0 {
				t.Errorf("got %d, want 0", componentsCount)
			}
			if componentsMask != 0 {
				t.Errorf("got %d, want 0", componentsMask)
			}
		})
	}
}
