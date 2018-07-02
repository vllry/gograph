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

func complexGraph1() Graph {
	g := NewGraph()
	v := make(map[string]*Vertex, 21)

	// Make vertices
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

/*func vertexIdsFromSlice(vertices []*Vertex) []string {
	ids := make([]string, len(vertices))
	for index, v := range vertices {
		ids[index] = v.id
	}
	return ids
}

func vertexIdsFromMap(vertices map[string]*Vertex) []string {
	ids := make([]string, len(vertices))
	for vid := range vertices {
		ids = append(ids, vid)
	}
	return ids
}*/
