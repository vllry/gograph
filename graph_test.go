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
