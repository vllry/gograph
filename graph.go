package gograph

type Graph struct {
	vertexSet map[string]*Vertex
}

func NewGraph() *Graph {
	return &Graph{
		vertexSet: make(map[string]*Vertex),
	}
}

func (g *Graph) copyVertex(v *Vertex) {
	u := v.copyWithoutAdjacency()
	g.vertexSet[u.id] = &u
}

func (g *Graph) Copy() *Graph {
	h := NewGraph()

	for _, v := range g.vertexSet {
		h.copyVertex(v)
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

func (g *Graph) Components() []*Graph {
	components := make([]*Graph, 0)
	inComponent := make(map[string]bool, 0)

	for _, firstVertexInComponent := range g.Vertices() {
		if inComponent[firstVertexInComponent.id] == true {
			continue
		}
		// Reaching here means firstVertexInComponent is in a to-be-discovered component.
		component := NewGraph() // Each loop happens once per component.

		// Perform a BFS until no more vertices are accessible. Add found vertices to the component.
		visitQueue := []*Vertex{firstVertexInComponent}
		for len(visitQueue) != 0 {
			current := visitQueue[0]
			visitQueue = visitQueue[1:]

			// Log a newly found vertex as being "in this component", update BFS accordingly.
			if _, present := component.Vertices()[current.id]; present == false {
				inComponent[current.id] = true
				component.copyVertex(current)
				for _, v := range current.GetAdjacent() {
					visitQueue = append(visitQueue, v)
				}
			}
		}

		// Now that all vertices are present, re-add all relevant edges from g.
		for _, v := range component.Vertices() {
			for _, adjInParent := range g.Vertices()[v.id].GetAdjacent() {
				adjInComponent, inComponent := component.Vertices()[adjInParent.id]
				if inComponent == true {
					e := v.Adjacent(adjInComponent)
					e.conformTo(adjInParent.Edges()[v.id])
				}
			}
		}

		components = append(components, component)
	}

	return components
}
