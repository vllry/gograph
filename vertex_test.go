package gograph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertVertex(t *testing.T) {
	g := NewGraph()
	v := g.Vertex("test")

	eq := reflect.DeepEqual(g.Vertices(), map[string]*Vertex{"test": v})
	if !eq {
		t.Error("Vertices in the graph did not match.")
	}
}

func TestVertexAttributes(t *testing.T) {
	key := "ping"
	value := "pong"

	g := NewGraph()
	v := g.Vertex("test")
	v.SetAttribute(key, value)

	actual, _ := v.GetAttribute(key)
	if actual != value {
		t.Error(fmt.Sprintf("Expected attribute %s:%s, got %s:%s", key, value, key, actual))
	}

	h := g.Copy()
	actual, _ = h.Vertices()["test"].GetAttribute(key)
	if actual != value {
		t.Error(fmt.Sprintf("Expected copied attribute %s:%s, got %s:%s", key, value, key, actual))
	}
}

func TestSimpleAdjacentVertices(t *testing.T) {
	g := NewGraph()
	v1 := g.Vertex("1")
	v2 := g.Vertex("2")

	v1.Adjacent(v2)

	eq := reflect.DeepEqual(v1.GetAdjacent(), map[string]*Vertex{"2": v2})
	if !eq {
		t.Error("v1 did not have expected Adjacent vertices")
	}
	eq = reflect.DeepEqual(v2.GetAdjacent(), map[string]*Vertex{"1": v1})
	if !eq {
		t.Error("v2 did not have expected Adjacent vertices")
	}
}

func TestComplexAdjacentVertices(t *testing.T) {
	g := NewGraph()

	v1 := g.Vertex("1")
	v2 := g.Vertex("2")
	v3 := g.Vertex("3")
	v4 := g.Vertex("4")

	v1.Adjacent(v2)
	v1.Adjacent(v3)
	v2.Adjacent(v3)
	v1.Adjacent(v4)

	var tests = []struct {
		vertex   *Vertex
		adjacent map[string]*Vertex
	}{
		{
			v1,
			map[string]*Vertex{
				"2": v2,
				"3": v3,
				"4": v4,
			},
		},
		{
			v2,
			map[string]*Vertex{
				"1": v1,
				"3": v3,
			},
		},
		{
			v3,
			map[string]*Vertex{
				"1": v1,
				"2": v2,
			},
		},
		{
			v4,
			map[string]*Vertex{
				"1": v1,
			},
		},
	}

	for _, test := range tests {
		eq := reflect.DeepEqual(test.vertex.GetAdjacent(), test.adjacent)
		if !eq {
			t.Error(fmt.Sprintf("%s did not have expected adjacent vertices", test.vertex.id))
		}
	}
}

func TestK5(t *testing.T) {
	k5 := k(5)
	ensureComplete(k5, t)
}

func ensureComplete(g *Graph, t *testing.T) {
	vertexMap := g.Vertices()

	for vid, v := range vertexMap {
		var expected = map[string]*Vertex{}
		for uid, u := range vertexMap {
			if vid != uid {
				expected[uid] = u
			} else {
			}
		}

		if !reflect.DeepEqual(v.GetAdjacent(), expected) {
			t.Error(fmt.Sprintf("%s didn't have expected adjacent vertices:", vid))
		}
	}
}
