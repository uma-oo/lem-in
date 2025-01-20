package parser

import (
	"fmt"
	"reflect"
)

// The idea behind is : to find the first group --> and then the second group, and keep only one of them based
// on the one who will minimise the number of turns and so on til reaching the last group
// This is used here just to be able to optimize the memory usage !!!

// there is one some edge cases where we'll be having the worst case of passing through all the possible solution (it's not anymore the edge case it's the rule)
// the case of 1 ANT , we need to favorise the shortest path of all the graph

func (G *Group) InitializeMvt(graph *Colony) []*Agent {
	var Agents map[int]*Agent = make(map[int]*Agent)
	var Agents_slice []*Agent
	shortest := G.Paths[0] // initial path
	is_first := false
	// This is the correct version of the function and nothing more or less

	for j := 1; j <= graph.Ants; j++ {
		if _, ok := Agents[j]; !ok {
			if !is_first {
				// First Ant to some inside the colony
				ant := NewAgent()
				ant.FindPath(j, graph, G, Agents)
				ant.CountPath(Agents_slice)
				// fmt.Printf("ant here inside : %v\n", ant)
				shortest.Length += 1
				is_first = true
				Agents[j] = ant
				Agents_slice = append(Agents_slice, ant)
			} else {
				shortest = G.ReturnShortest()
				new_ant := NewAgent()
				new_ant.FindPath(j, graph, G, Agents)
				new_ant.CountPath(Agents_slice)
				shortest.Length += 1
				Agents[j] = new_ant
				Agents_slice = append(Agents_slice, new_ant)

			}
		}
	}

	return Agents_slice
}

func (g *Group) MoveAnts(graph *Colony) []string {
	lines := []string{}
	agents := g.InitializeMvt(graph) // Find the paths for each ant in the colony

	for i := 1; i <= g.Turns; i++ {

		line := ""
		positions := make(map[string]struct{})
		for j, ant := range agents {
			pos := GetAntPos(i, ant.Pos, ant.PathUsed.Rooms_found)
			_, ok := positions[pos]
			// fmt.Printf("ant: %v pos: %v turn: %v\n", ant, pos, i)

			if pos != "" && !ok {
				line += fmt.Sprintf("L%v-%v ", j+1, pos)
				if pos != graph.End_room.Name {
					positions[pos] = struct{}{}
				}

			}

		}
		fmt.Println(line)
	}
	return lines
}

// given a specific ant and turn we get l position dyal ant f l path
func GetAntPos(turn int, ant int, path []string) string {
	if turn-ant < len(path) && turn-ant >= 0 {
		return path[turn-ant]
	}
	return ""
}

// see after if we can edit this with the turn as a parameter
func (A *Agent) FindPath(ant int, graph *Colony, group_chosen *Group, agents map[int]*Agent) {
	// Find The Shortest Path inside The group and then assign it to the ant
	shortest_path := group_chosen.ReturnShortest()
	A.PathUsed = shortest_path
}

// The idea is as follows
// if the path has been taken by another ant meaning by this, it's not her first time to appear
// we index the ant using the Pos and the Pos means only that it's the first the second or the third in the path
// like having a primary key (path , pos) pos reflects the turn for that specific path

func (A *Agent) CountPath(agents []*Agent) {
	count := 1
	for _, agent := range agents {
		if reflect.DeepEqual(A.PathUsed.Rooms_found, agent.PathUsed.Rooms_found) {
			count++
		}
	}
	A.Pos = count
}
