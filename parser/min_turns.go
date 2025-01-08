package parser

// THIS WILL BE USED IN A REAL TIME CHECKING OF SHOULD THE GROUP BE KEPT OR NO!!
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
				shortest = G.ReturnShortest()
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

// Used in the turns calcul
// Returns the shortest path always
func (G *Group) ReturnShortest() *Path {
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

//******************************************************************************************//

// Not good for memory what if we have 10000 group
func DecideWhichGroup(graph *Colony) *Group {
	groups := FindAllGroups(graph)
	chosen_group := groups[0]
	groups[0].CalculTurns(graph)
	min_turns := groups[0].Turns
	for _, g := range groups {
		g.CalculTurns(graph)
		if min_turns < g.Turns {
			min_turns = g.Turns
			chosen_group = g
		}
	}
	return chosen_group
}
