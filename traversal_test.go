package gograph

import (
	"fmt"
	"testing"
)

func TestPathExists(t *testing.T) {
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
