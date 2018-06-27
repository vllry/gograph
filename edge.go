package gograph

type Edge struct {
	attributes map[string]string
}

func NewEdge() Edge {
	return Edge{
		attributes: make(map[string]string),
	}
}
