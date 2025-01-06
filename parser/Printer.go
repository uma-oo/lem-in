package parser

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
}

// func (G *Group) CalculTurns(graph *Colony) int {
// 	for i := 1; i <= graph.Ants; i++ {
// 		for _ , path := range G.Paths {
           
// 		}
// 	}
// 	return 0
// }
