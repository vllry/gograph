package gograph

/*
BFS implementation.
O(|V(G)| + |V(E)|)
Slightly cheaper than Dijkstra's algorithm when simply checking for connectivity.
*/
func (g *Graph) PathExists(u *Vertex, v *Vertex) bool {
	visitedMap := make(map[string]*Vertex, 0)
	queue := []*Vertex{u}

	for len(queue) != 0 {
		// Don't double queue
		// Still having trouble queueing v5
		current := queue[0]
		queue = queue[1:]
		if visitedMap[current.id] != nil {
			continue // Skip if already seen before - slight optimization.
			// Already-checked vertices are not re-queued.
		}

		if current.id == v.id {
			return true
		}
		visitedMap[current.id] = current

		for _, w := range current.GetAdjacent() {
			if visitedMap[w.id] == nil {
				queue = append(queue, w)
			}
		}

	}

	return false
}

//func (g *Graph) CheaptestPath(costAttribute string, defaultCost int) {
//}
