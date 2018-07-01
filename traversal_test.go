package gograph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShortestPath(t *testing.T) {
	g := complexGraph1()
	vertices := g.Vertices()

	path := g.ShortestPath(vertices["0"], vertices["20"])
	pathIds := make([]string, len(path))
	for _, v := range path {
		pathIds = append(pathIds, v.id)
	}

	expected := []string{"0", "20"}
	if !reflect.DeepEqual(pathIds, expected) {
		t.Error(fmt.Sprintf("Got incorrect shortest path: %s", pathIds))
	}
}
