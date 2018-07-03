package gograph

import "fmt"

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

/*
Dijkstra's algorithm - goto for cheapest paths.
https://en.wikipedia.org/wiki/Dijkstra's_algorithm
*/
func (g *Graph) ShortestPath(start *Vertex, target *Vertex) []*Vertex {
	maxDistance := int(^uint(0) >> 1) // Max int
	queue := make([]*Vertex, 0)
	prev := make(map[string]*Vertex, 0)
	distances := make(map[string]int)

	// Initialize tracking info for vertices.
	for _, v := range g.Vertices() {
		prev[v.id] = nil
		distances[v.id] = maxDistance
		queue = append(queue, v)
	}
	distances[start.id] = 0

	for len(queue) != 0 {
		// Find the closest vertex in the queue.
		var closest *Vertex
		var closestIndex int
		for index, v := range queue {
			if closest == nil || distances[v.id] < distances[closest.id] {
				closest = v
				closestIndex = index
			}
		}
		fmt.Println("dequeueing", closest.id)
		queue = append(queue[:closestIndex], queue[closestIndex+1:]...)

		for _, v := range closest.GetAdjacent() {
			alt := distances[closest.id] + 1 // +1 is the cost of the edge
			if alt < distances[v.id] {
				distances[v.id] = alt
				prev[v.id] = closest
			}
		}
	}

	fmt.Println("path assembly")

	// Re-assemble cheapest reversePath from cost data.
	var reversePath []*Vertex
	tail := target
	for tail.id != start.id {
		fmt.Println("reverse path", tail.id)
		reversePath = append(reversePath, tail)
		tail = prev[tail.id] // Move closer to the source
	}
	fmt.Println("Path reversal")
	path := make([]*Vertex, len(reversePath))
	for index, v := range reversePath {
		fmt.Println(v.id)
		path[len(reversePath)-1-index] = v
	}

	return reversePath
}
