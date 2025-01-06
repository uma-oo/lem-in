package parser

import "fmt"

// The idea behind is : to finc the first group --> and then the second group, and keep only one of them based
// on the one who will minimise the number of turns and so on til reaching the last group
// This is used here just to be able to optimize the memory usage !!!

// there is one some edge cases where we'll be having the worst case of passing through all the possible solution
// the case of 1 ANT , we need to favorise the shortest path of all the graph

func PrintTurns(paths [][]string) {
}

func TurnsWithShortest(graph *Colony, group *Group) int {
	return (len(group.Paths) + graph.Ants) - 1
}

func DecideWhichGroup(graph *Colony) {
	groups:=FindAllGroups(graph)
		for _, g := range groups {
			fmt.Printf("g.CalculTurns(graph): %v\n", g.CalculTurns(graph))
		}
	
}

func (G *Group) CalculTurns(graph *Colony) int {
	shortest_path := G.Paths[0]
	for i := 1; i <= graph.Ants; i++ {
		for _, path := range G.Paths {
			if path.Length < shortest_path.Length {
				
				shortest_path=path
				path.Length += 1
			} else {
			
				shortest_path.Length+=1

			
			}
		}
	}
	return shortest_path.Length
}
