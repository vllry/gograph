package gograph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGraph_PathExists(t *testing.T) {
	g := complexGraph1()
	vertices := g.Vertices()

	path := g.PathExists(vertices["0"], vertices["20"])
	if !path {
		t.Error(fmt.Sprintf("Expecting to find a path from 0 to 20"))
	}

	path = g.PathExists(vertices["0"], vertices["12"])
	if path {
		t.Error(fmt.Sprintf("Not expecting to find a path from 0 to 13"))
	}
}

func TestGraph_ShortestPath(t *testing.T) {
	h := complexGraph1()
	g := h.Copy()

	v := g.Vertices()

	var path []*Vertex
	var cost int

	path, cost = g.ShortestPath(v["0"], v["20"])
	expectedPath := []*Vertex{v["0"], v["3"], v["5"], v["14"], v["16"], v["19"], v["20"]}
	expectedCost := 6
	if !reflect.DeepEqual(path, expectedPath) {
		t.Error(fmt.Sprintf(
			"Failed to find expected path from 0 to 20. Wanted %s, got %s",
			pathToString(expectedPath),
			pathToString(path),
		))
	}
	if cost != expectedCost {
		t.Error(fmt.Sprintf("Expected cost %d, got %d", expectedCost, cost))
	}

	path, cost = g.ShortestPath(v["0"], v["12"])
	expectedCost = int(^uint(0) >> 1)
	if len(path) != 0 {
		t.Error(fmt.Sprintf("Expected a path length of 0"))
	}
	if cost != expectedCost {
		t.Error(fmt.Sprintf("Expected cost %d, got %d", expectedCost, cost))
	}
}

func TestGraph_CheapestPath(t *testing.T) {
	g := complexGraph2()
	v := g.Vertices()

	var path []*Vertex
	var cost int

	path, cost = g.CheapestPath(v["4"], v["17"], "cost", 1)
	expectedPath := []*Vertex{v["4"], v["3"], v["0"], v["7"], v["8"], v["9"], v["12"], v["18"], v["17"]}
	expectedCost := 11
	if !reflect.DeepEqual(path, expectedPath) {
		t.Error(fmt.Sprintf(
			"Failed to find expected path from 0 to 20. Wanted %s, got %s",
			pathToString(expectedPath),
			pathToString(path),
		))
	}
	if cost != expectedCost {
		t.Error(fmt.Sprintf("Expected cost %d, got %d", expectedCost, cost))
	}
}

func pathToString(s []*Vertex) string {
	str := ""
	for _, v := range s {
		str += v.id + " "
	}
	return str
}
