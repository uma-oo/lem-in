package parser

import "fmt"

// The idea behind is : to finc the first group --> and then the second group, and keep only one of them based
// on the one who will minimise the number of turns and so on til reaching the last group
// This is used here just to be able to optimize the memory usage !!!

// there is one some edge cases where we'll be having the worst case of passing through all the possible solution
// the case of 1 ANT , we need to favorise the shortest path of all the graph

// Work on this before getting to the leak of memory problem

// We need to make sure that each room in the path is not used twice in the same path
func (G *Group) InitializeMvt(graph *Colony) {
	for i := 0; i < G.Turns; i++ { // steps
		for j := 1; j <= graph.Ants; j++ {
			agent := NewAgent()
			shortest := G.ReturnShortest()
			agent.PathUsed = shortest
			for _, room := range agent.PathUsed.Rooms_found {
				fmt.Printf("room: %v\n", room)
			}

		}
	}
}





// given a specific ant and turn we get l position dyal ant f l path 
func GetAntPos(turn int, ant int, path []string) string {
	if turn-ant >= 0 {
		return path[turn-ant]
	}
	return ""
}



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
