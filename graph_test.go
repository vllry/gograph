package gograph

import (
	"fmt"
	"testing"
)

func TestNewGraph(t *testing.T) {
	_ = NewGraph()
}

func TestBasicCopy(t *testing.T) {
	g := k(4)
	h := g.Copy()

	expectedOrder := 4
	actualOrder := h.Order()

	if expectedOrder != actualOrder {
		t.Error(fmt.Sprintf("Graph was order %d, expected %d", actualOrder, expectedOrder))
	}
	ensureComplete(h, t)
}

func TestGraph_Components(t *testing.T) {
	g := complexGraph2()
	components := g.Components()

	expectedComponents := 2
	actualComponents := len(components)
	if expectedComponents != actualComponents {
		t.Error(fmt.Sprintf("Expected %d components, got %d", expectedComponents, actualComponents))
	}

	componentA, componentB := components[0], components[1]
	smallComponent := componentA
	bigComponent := componentB
	if componentB.Order() < componentA.Order() {
		smallComponent = componentB
		bigComponent = componentA
	}

	if smallComponent.Order() != 4 || bigComponent.Order() != 19 {
		t.Error(fmt.Sprintf("Expected component sizes of %d and %d - got %d and %d",
			4, 19, smallComponent.Order(), bigComponent.Order()))
	}

	// Test on k1.
	k := k(1)
	components = k.Components()
	expectedComponents = 1
	actualComponents = len(components)
	if expectedComponents != actualComponents {
		t.Error(fmt.Sprintf("Expected %d components, got %d", expectedComponents, actualComponents))
	}
}
