package gograph

import (
	"reflect"
	"testing"
	"fmt"
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

func TestAttributes(t *testing.T) {
	key := "ping"
	value := "pong"

	v := NewVertex("attr test")
	v.SetAttribute(key, value)

	actual, _ := v.GetAttribute(key)
	if actual != value {
		t.Error(fmt.Sprintf("Expected attributed %s:%s, got %s:%s", key, value, key, actual))
	}
}

func TestSimpleAdjacentVertices(t *testing.T) {
	v1 := NewVertex("1")
	v2 := NewVertex("2")

	v1.Adjacent(&v2)

	eq := reflect.DeepEqual(v1.GetAdjacent(), map[string]*Vertex{"2": &v2})
	if !eq {
		t.Error("v1 did not have expected Adjacent vertices")
	}
	eq = reflect.DeepEqual(v2.GetAdjacent(), map[string]*Vertex{"1": &v1})
	if !eq {
		t.Error("v2 did not have expected Adjacent vertices")
	}
}

func TestComplexAdjacentVertices(t *testing.T) {
	v1 := NewVertex("1")
	v2 := NewVertex("2")
	v3 := NewVertex("3")
	v4 := NewVertex("4")

	v1.Adjacent(&v2)
	v1.Adjacent(&v3)
	v2.Adjacent(&v3)
	v1.Adjacent(&v4)

	var tests = []struct {
		vertex  *Vertex
		adjacent map[string]*Vertex
	}{
		{
			&v1,
			map[string]*Vertex{
				"2": &v2,
				"3": &v3,
				"4": &v4,
			},
		},
		{
			&v2,
			map[string]*Vertex{
				"1": &v1,
				"3": &v3,
			},
		},
		{
			&v3,
			map[string]*Vertex{
				"1": &v1,
				"2": &v2,
			},
		},
		{
			&v4,
			map[string]*Vertex{
				"1": &v1,
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