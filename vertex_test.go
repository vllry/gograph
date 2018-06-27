package gograph

import (
	"reflect"
	"testing"
)

func TestNewVertex(t *testing.T) {
	_ = NewVertex("test")
}

func TestInsertVertex(t *testing.T) {
	g := NewGraph()
	v := NewVertex("test")

	g.add(&v)
	eq := reflect.DeepEqual(g.vertices(), map[string]*Vertex{"test": &v})
	if !eq {
		t.Error("Vertices in the graph did not match.")
	}
}

func TestAdjacentVertices(t *testing.T) {
	v1 := NewVertex("1")
	v2 := NewVertex("2")
	v1.adjacent(&v2)

	eq := reflect.DeepEqual(v1.getAdjacent(), map[string]*Vertex{"2": &v2})
	if !eq {
		t.Error("v1 did not have expected adjacent vertices")
	}

	eq = reflect.DeepEqual(v2.getAdjacent(), map[string]*Vertex{"1": &v1})
	if !eq {
		t.Error("v2 did not have expected adjacent vertices")
	}
}
