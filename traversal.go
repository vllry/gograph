package gograph

import "fmt"

/*
BFS implementation.
O(|V(G)| + |V(E)|)
*/
func (g *Graph) ShortestPath(u *Vertex, v *Vertex) []*Vertex {
	visitedMap := make(map[string]*Vertex, 0)
	visitedList := make([]*Vertex, 0)
	queue := []*Vertex{u}

	for len(queue) != 0 {
		// Don't double queue
		// Still having trouble running for v5
		current := queue[0]
		queue = queue[1:]
		fmt.Println(current.id, "adj", current.GetAdjacentIds())

		visitedList = append(visitedList, current)
		if current.id == v.id {
			return visitedList
		}
		visitedMap[current.id] = current

		for m := range visitedMap {
			fmt.Println(m)
		}
		for _, w := range u.GetAdjacent() {
			if visitedMap[w.id] == nil {
				fmt.Println("queuing", w.id)
				queue = append(queue, w)
			}
		}
		fmt.Println("q", queue)
		fmt.Println("visited", visitedList)
	}

	return make([]*Vertex, 0)
}
