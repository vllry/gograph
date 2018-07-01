package gograph

type Edge struct {
	attributes    map[string]string
	verticesMap   map[string]*Vertex
	verticesSlice [2]*Vertex
}

func newEdge(v1 *Vertex, v2 *Vertex) Edge {
	return Edge{
		attributes: make(map[string]string),
		verticesMap: map[string]*Vertex{
			v1.id: v1,
			v2.id: v2,
		},
		verticesSlice: [2]*Vertex{
			v1,
			v2,
		},
	}
}

func (e *Edge) GetIncident() (*Vertex, *Vertex) {
	return e.verticesSlice[0], e.verticesSlice[1]
}
