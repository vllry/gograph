package gograph

type Edge struct {
	Attributes map[string]string
	vertices map[string]*Vertex
}

func NewEdge(v1 *Vertex, v2 *Vertex) Edge {
	return Edge{
		Attributes: make(map[string]string),
		vertices:   map[string]*Vertex{v1.id: v1, v2.id:v2},
	}
}
