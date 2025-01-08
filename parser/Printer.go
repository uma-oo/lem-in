package parser

// The idea behind is : to find the first group --> and then the second group, and keep only one of them based
// on the one who will minimise the number of turns and so on til reaching the last group
// This is used here just to be able to optimize the memory usage !!!

// there is one some edge cases where we'll be having the worst case of passing through all the possible solution
// the case of 1 ANT , we need to favorise the shortest path of all the graph

// Work on this before getting to the leak of memory problem

// We need to make sure that each room in the path is not used twice in the same path
// Let's do it for one ant for now

func (G *Group) InitializeMvt(graph *Colony) map[int]*Agent {
	var Agents map[int]*Agent = make(map[int]*Agent)
	shortest := G.Paths[0] // initial path
	is_first := false
	for i := 1; i <= G.Turns; i++ { // steps
		is_filled := false
		if !is_filled {
			for j := 1; j <= graph.Ants; j++ {
				if _, ok := Agents[j]; !ok {
					if !is_first {
						// First Ant to some inside the colony
						ant := NewAgent()
						ant.FindPath(j, graph, G, i)
						shortest.Length += 1
						is_first = true
						Agents[j] = ant
                        
					} else {
						shortest = G.ReturnShortest()
						new_ant := NewAgent()
						new_ant.FindPath(j, graph, G, i)
						shortest.Length += 1
						Agents[j] = new_ant

					}
					is_filled = true

				}
			}
		}

	}
	return Agents
}

// given a specific ant and turn we get l position dyal ant f l path
func GetAntPos(turn int, ant int, path []string) string {
	if turn-ant >= 0 {
		return path[turn-ant]
	}
	return ""
}

// see after if we can edit this with the turn as a parameter
func (A *Agent) FindPath(ant int, graph *Colony, group_chosen *Group, turn int) {
	// Find The Shortest Path inside The group and then assign it to the ant

	shortest_path := group_chosen.ReturnShortest()
	A.Name = ant
	A.PathUsed = shortest_path
	ant_pos := turn-A.Name
	if ant_pos >= 0 {
		A.Start_Step = ant_pos
	}
	A.End_Step = len(A.PathUsed.Rooms_found) - A.Start_Step - 1 // to know where we need to stop printing
}
