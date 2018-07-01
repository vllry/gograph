package gograph

import (
	"strconv"
)

func k(order int) Graph {
	g := NewGraph()

	var vertices = make([]*Vertex, 0)
	for i := 0; i < order; i++ {
		id := strconv.Itoa(i)
		v := g.Vertex(id)
		for _, u := range vertices {
			v.Adjacent(u)
		}
		vertices = append(vertices, v)
	}

	return g
}
