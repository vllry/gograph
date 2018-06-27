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
		t.Fail()
	}
}
