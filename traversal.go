package gograph

/*
BFS implementation.
O(|V(G)| + |V(E)|) time
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

/*
Dijkstra's algorithm - goto for cheapest paths.
https://en.wikipedia.org/wiki/Dijkstra's_algorithm
*/
func (g *Graph) ShortestPath(start *Vertex, target *Vertex) ([]*Vertex, int) {
	maxDistance := int(^uint(0) >> 1) // Max int
	queue := make([]*Vertex, 0)
	prev := make(map[string]*Vertex, 0)
	costs := make(map[string]int)

	// Initialize tracking info for vertices.
	for _, v := range g.Vertices() {
		prev[v.id] = nil
		costs[v.id] = maxDistance
		queue = append(queue, v)
	}
	costs[start.id] = 0

	for len(queue) != 0 {
		// Find the closest vertex in the queue.
		// A priority queue would be faster than a for-loop.
		var closest *Vertex
		var closestIndex int
		for index, v := range queue {
			if closest == nil || costs[v.id] < costs[closest.id] {
				closest = v
				closestIndex = index
			}
		}
		queue = append(queue[:closestIndex], queue[closestIndex+1:]...) // Remove that element from the queue.

		for _, v := range closest.GetAdjacent() {
			alt := costs[closest.id] + 1 // +1 is the cost of the edge
			if alt < costs[v.id] {
				costs[v.id] = alt
				prev[v.id] = closest
			}
		}
	}

	// Re-assemble cheapest reversePath from cost data.
	var reversePath []*Vertex
	tail := target
	for prev[tail.id] != nil {
		reversePath = append(reversePath, tail)
		tail = prev[tail.id] // Move closer to the source
	}
	if len(reversePath) != 0 { // Loop returns before reaching the start vetex.
		reversePath = append(reversePath, start)
	}

	// Reverse the path to reflect traversal order.
	path := make([]*Vertex, len(reversePath))
	for index, v := range reversePath {
		path[len(reversePath)-1-index] = v
	}

	cheaptestCost := costs[target.id]

	return reversePath, cheaptestCost
}
