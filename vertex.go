package gograph

type Vertex struct {
	id         string
	attributes map[string]string
	adjacentTo map[string]*Vertex
	edges      map[string]*Edge
}

func (v *Vertex) Adjacent(v2 *Vertex) *Edge {
	v.adjacentTo[v2.id] = v2
	v2.adjacentTo[v.id] = v

	e := newEdge(v, v2)
	v.edges[v2.id] = &e
	v2.edges[v.id] = &e

	return &e
}

// Copies a vertex, without adjacency information.
func (v *Vertex) copyWithoutAdjacency() Vertex {
	u := newVertex(v.id)

	u.attributes = v.attributes

	return u
}

func newVertex(id string) Vertex {
	v := Vertex{
		id:         id,
		attributes: make(map[string]string),
		adjacentTo: make(map[string]*Vertex),
		edges:      make(map[string]*Edge),
	}
	return v
}

func (v *Vertex) GetAdjacent() map[string]*Vertex {
	return v.adjacentTo
}

func (v *Vertex) GetAdjacentIds() []string {
	ids := make([]string, 0)
	for id := range v.GetAdjacent() {
		ids = append(ids, id)
	}
	return ids
}

func (v *Vertex) SetAttribute(key string, value string) {
	v.attributes[key] = value
}

func (v *Vertex) GetAttribute(key string) (string, error) {
	value := v.attributes[key] // TODO handle errors
	return value, nil
}
