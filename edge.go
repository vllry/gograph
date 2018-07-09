package gograph

type Edge struct {
	intAttributes    map[string]int
	stringAttributes map[string]string
	verticesMap      map[string]*Vertex
	verticesSlice    [2]*Vertex
}

func newEdge(v1 *Vertex, v2 *Vertex) Edge {
	return Edge{
		intAttributes:    make(map[string]int),
		stringAttributes: make(map[string]string),
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

func (e *Edge) conformTo(e2 *Edge) {
	e.intAttributes = e2.intAttributes
	e.stringAttributes = e2.stringAttributes
}

func (e *Edge) GetIncident() (*Vertex, *Vertex) {
	return e.verticesSlice[0], e.verticesSlice[1]
}

func (e *Edge) SetIntAttribute(key string, value int) {
	e.intAttributes[key] = value
}

func (e *Edge) GetIntAttribute(key string) (int, error) {
	value := e.intAttributes[key]
	return value, nil
}

func (e *Edge) SetStringAttribute(key string, value string) {
	e.stringAttributes[key] = value
}

func (e *Edge) GetStringAttribute(key string) (string, error) {
	value := e.stringAttributes[key]
	return value, nil
}
