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
