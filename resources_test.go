package gograph

import (
	"strconv"
)

func k(order int) *Graph {
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

// Creates a graph with 1 "main" component and several pendant vertices.
func complexGraph1() *Graph {
	g := NewGraph()
	v := make(map[string]*Vertex, 21)

	i := 0
	for i < 21 {
		u := g.Vertex(strconv.Itoa(i))
		v[u.id] = u
		i += 1
	}

	v["0"].Adjacent(v["1"])
	v["0"].Adjacent(v["2"])
	v["0"].Adjacent(v["3"])
	v["1"].Adjacent(v["4"])
	v["2"].Adjacent(v["4"])
	v["4"].Adjacent(v["5"])
	v["3"].Adjacent(v["5"])
	v["5"].Adjacent(v["14"])
	v["14"].Adjacent(v["15"])
	v["14"].Adjacent(v["16"])
	v["15"].Adjacent(v["19"])
	v["16"].Adjacent(v["19"])
	v["17"].Adjacent(v["19"])
	v["18"].Adjacent(v["19"])
	v["19"].Adjacent(v["20"])

	return g
}

// Creates a graph with weights and 2 components.
func complexGraph2() *Graph {
	g := NewGraph()
	v := make(map[string]*Vertex, 21)

	i := 0
	for i < 23 {
		u := g.Vertex(strconv.Itoa(i))
		v[u.id] = u
		i += 1
	}

	v["0"].Adjacent(v["3"])
	v["0"].Adjacent(v["7"])
	v["3"].Adjacent(v["2"])
	v["3"].Adjacent(v["4"])
	v["2"].Adjacent(v["1"]).SetIntAttribute("cost", 7)
	v["2"].Adjacent(v["5"]).SetIntAttribute("cost", 5)
	v["4"].Adjacent(v["6"])
	v["5"].Adjacent(v["6"])
	v["6"].Adjacent(v["8"]).SetIntAttribute("cost", 12)
	v["8"].Adjacent(v["7"])
	v["8"].Adjacent(v["9"]).SetIntAttribute("cost", 3)
	v["9"].Adjacent(v["10"])
	v["9"].Adjacent(v["11"]).SetIntAttribute("cost", 3)
	v["9"].Adjacent(v["12"])
	v["11"].Adjacent(v["15"])
	v["12"].Adjacent(v["13"]).SetIntAttribute("cost", 2)
	v["12"].Adjacent(v["15"]).SetIntAttribute("cost", 4)
	v["12"].Adjacent(v["18"]).SetIntAttribute("cost", 2)
	v["13"].Adjacent(v["14"])
	v["13"].Adjacent(v["17"]).SetIntAttribute("cost", 4)
	v["15"].Adjacent(v["16"])
	v["17"].Adjacent(v["18"])

	v["19"].Adjacent(v["20"]).SetIntAttribute("cost", 4)
	v["20"].Adjacent(v["21"])
	v["21"].Adjacent(v["22"])
	v["22"].Adjacent(v["19"])

	return g
}
