package gograph

import (
	"fmt"
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

func TestEdgeAttributes(t *testing.T) {
	g := NewGraph()
	v1 := g.Vertex("1")
	v2 := g.Vertex("2")
	e := v1.Adjacent(v2)

	attr1 := 12
	attr2 := 2
	attr3 := "test"

	e.SetIntAttribute("cost", attr1)
	e.SetIntAttribute("size", attr2)
	e.SetStringAttribute("type", attr3)

	key := "cost"
	actual, _ := e.GetIntAttribute(key)
	expected := attr1
	if actual != expected {
		t.Error(fmt.Sprintf("Expected attribute %s:%d, got %s:%d", key, expected, key, actual))
	}

	key = "size"
	actual, _ = e.GetIntAttribute(key)
	expected = attr2
	if actual != expected {
		t.Error(fmt.Sprintf("Expected attribute %s:%d, got %s:%d", key, expected, key, actual))
	}

	key = "type"
	actualString, _ := e.GetStringAttribute(key)
	expectedString := attr3
	if actual != expected {
		t.Error(fmt.Sprintf("Expected attribute %s:%s, got %s:%s", key, expectedString, key, actualString))
	}

	h := g.Copy()
	e2 := h.Vertices()["1"].edges["2"]

	key = "cost"
	actual, _ = e2.GetIntAttribute(key)
	expected = attr1
	if actual != expected {
		t.Error(fmt.Sprintf("Expected attribute %s:%d, got %s:%d", key, expected, key, actual))
	}

	key = "size"
	actual, _ = e2.GetIntAttribute(key)
	expected = attr2
	if actual != expected {
		t.Error(fmt.Sprintf("Expected attribute %s:%d, got %s:%d", key, expected, key, actual))
	}

	key = "type"
	actualString, _ = e2.GetStringAttribute(key)
	expectedString = attr3
	if actual != expected {
		t.Error(fmt.Sprintf("Expected attribute %s:%s, got %s:%s", key, expectedString, key, actualString))
	}
}
