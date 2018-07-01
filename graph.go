package gograph

type Graph struct {
	vertexSet map[string]*Vertex
}

func NewGraph() Graph {
	return Graph{
		vertexSet: make(map[string]*Vertex),
	}
}

func (g *Graph) Vertex(id string) *Vertex {
	v := Vertex{
		id:         id,
		attributes: make(map[string]string),
		adjacentTo: make(map[string]*Vertex),
		edges:      make(map[string]*Edge),
	}
	g.vertexSet[v.id] = &v

	return &v
}

func (g *Graph) Vertices() map[string]*Vertex {
	return g.vertexSet
}

func (g *Graph) VertexIds() []string {
	ids := make([]string, 0)
	for id := range g.Vertices() {
		ids = append(ids, id)
	}
	return ids
}
