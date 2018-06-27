package gograph

type Vertex struct {
	id         string
	adjacentTo map[string]*Vertex
}

func NewVertex(id string) Vertex {
	return Vertex{
		id:         id,
		adjacentTo: make(map[string]*Vertex),
	}
}

func (v1 *Vertex) adjacent(v2 *Vertex) {
	v1.adjacentTo[v2.id] = v2
}
