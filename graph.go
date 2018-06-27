package gograph

type Graph struct {
	vertexSet map[string]*Vertex
}

func NewGraph() Graph {
	return Graph{
		vertexSet: make(map[string]*Vertex),
	}
}

func (g *Graph) add(v *Vertex) {
	g.vertexSet[v.id] = v
}

func (g *Graph) vertices() map[string]*Vertex {
	return g.vertexSet
}
