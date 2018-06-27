package gograph

type Vertex struct {
	id         string
	attributes map[string]string
	adjacentTo map[string]*Vertex
	edges      map[string]*Edge
}

func NewVertex(id string) Vertex {
	return Vertex{
		id:         id,
		attributes: make(map[string]string),
		adjacentTo: make(map[string]*Vertex),
		edges:      make(map[string]*Edge),
	}
}

func (v1 *Vertex) adjacent(v2 *Vertex) {
	v1.adjacentTo[v2.id] = v2
	v2.adjacentTo[v1.id] = v1

	e := NewEdge()
	v1.edges[v2.id] = &e
	v2.edges[v1.id] = &e
}

func (v *Vertex) getAdjacent() map[string]*Vertex {
	return v.adjacentTo
}
