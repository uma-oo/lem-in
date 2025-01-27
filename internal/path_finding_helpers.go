package internal

func MarkRoomsVisited(whole map[string]struct{}, part []string) {
	for _, element := range part {
		whole[element] = struct{}{}
	}
}

func Compare2Groups(graph *Colony, group1 *Group, group2 *Group) *Group {
	if group1.Turns <= group2.Turns {
		return group1
	}
	return group2
}

// The function that calculates turns needed for each group
func (G *Group) CalculTurns(graph *Colony) {
	shortest_path := G.Paths[0]
	shortest := shortest_path
	for i := 1; i <= graph.Ants; i++ {
		moved := false
		if i == 1 {
			shortest_path.Length += 1
			shortest = shortest_path
			moved = true
		} else {
			if !moved {
				shortest = G.ReturnShortestPath()
				shortest.Length += 1
			}
		}
	}
	G.Turns = shortest_path.Length - 1
	// Reset the length of the Paths
	for _, path := range G.Paths {
		path.Length = len(path.Rooms_found)
	}
}

func (G *Group) ReturnShortestPath() *Path {
	min := G.Paths[0].Length
	shortest := G.Paths[0]
	for _, p := range G.Paths {
		if p.Length < min {
			shortest = p
			min = shortest.Length
		}
	}
	return shortest
}
