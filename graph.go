package gograph

type Graph struct {
	vertexSet map[string]*Vertex
}

func NewGraph() *Graph {
	return &Graph{
		vertexSet: make(map[string]*Vertex),
	}
}

func (g *Graph) Copy() *Graph {
	h := NewGraph()

	// Copy vertices.
	for _, v := range g.vertexSet {
		u := v.copyWithoutAdjacency()
		h.vertexSet[u.id] = &u
	}

	// Link vertices.
	for _, v := range g.vertexSet {
		for _, vAdj := range v.GetAdjacent() {
			if h.vertexSet[v.id].GetAdjacent()[vAdj.id] == nil { // Only link if not already linked.
				h.vertexSet[v.id].Adjacent(
					h.vertexSet[vAdj.id],
				)
				h.vertexSet[v.id].edges[vAdj.id].conformTo(
					v.edges[vAdj.id],
				)
			}
		}
	}

	return h
}

func (g *Graph) Vertex(id string) *Vertex {
	v := newVertex(id)
	g.vertexSet[v.id] = &v
	return &v
}

func (g *Graph) Vertices() map[string]*Vertex {
	return g.vertexSet
}

func (g *Graph) Order() int {
	return len(g.vertexSet)
}

func (g *Graph) VertexIds() []string {
	ids := make([]string, 0)
	for id := range g.Vertices() {
		ids = append(ids, id)
	}
	return ids
}
