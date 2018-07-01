package gograph

import (
	"testing"
)

func TestNewEdge(t *testing.T) {
	g := NewGraph()
	v1 := g.Vertex("1")
	v2 := g.Vertex("2")
	e := v1.Adjacent(v2)

	u1, u2 := e.GetIncident()
	if u1 != v1 || u2 != v2 {
		t.Error("Edge between v1,v2 failed to list those vertices.")
	}
}
