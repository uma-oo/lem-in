package parser

import (
	"fmt"
	"runtime"
	"slices"
)

// used BFS hna to determine levels
// this is the implementation based on this stackoverflow question
// link hahuwa for more details : https://stackoverflow.com/questions/14144071/finding-all-the-shortest-paths-between-two-nodes-in-unweighted-undirected-graph?rq=1
// BFS + reverse DFS
// This is the part where the memory is leaking
// Problem detected is in the slice of strings that adds the same parent more than once
func Levels(graph *Colony, start string, end string) map[string]int {
	level := make(map[string]int)
	var current string
	var traversal *Traversal = NewTraversal() // Initilaize dakshi lkula traversal
	traversal.Is_Visited[graph.Start_room.Name] = true
	traversal.Visited_Node[start] = []string{}               // set that way to be known as the base case where the program will stop
	start_element := SetNode(start)                          // kaykunu 3ndna string khasshum yt7wlu t structs li 3ndna
	traversal.Queue = append(traversal.Queue, start_element) // appendiw element lwlani l queue
	level[start] = 1
	level[graph.Start_room.Name] = 0
   
	for len(traversal.Queue) > 0 {
		current = traversal.Pop()
		if current == end {
			break
		}

		for element := range graph.Tunnels[current].Links {
			if _, ok := traversal.Visited_Node[element]; !ok {
				traversal.Is_Visited[current] = true
				if !slices.Contains(traversal.Visited_Node[element], current) {
					traversal.Visited_Node[element] = append(traversal.Visited_Node[element], current)
				}

				level[element] = level[current] + 1
				
				node_element := SetNode(element)
				traversal.Queue = append(traversal.Queue, node_element)
				// printAlloc()

			} else {
				for _, parent := range traversal.Visited_Node[element] {
					if SameLevel(parent, current, level) && !slices.Contains(traversal.Visited_Node[element], current) {
						traversal.Visited_Node[element] = append(traversal.Visited_Node[element], current)
					}
				}
			}
		}
	}
	// fmt.Println("level",level)
	// fmt.Println(traversal.String())

	return level
}

// Backtracking bash n3awdu njbduu all the paths based on the levels we already have
// we can have more than one shortest path
// does not work well for all the cases somehow for the G0 case it fails to find the path G0-C0....
// This is my be due to the fact this path is not the shortest at all and so we will be needing to use the BFS in this case

func DFS(graph *Colony, start string) [][]string {
	var trajectories [][]string
	visited := make(map[string]bool)
	visited[graph.Start_room.Name] = true
	levels := Levels(graph, start, graph.End_room.Name)
	fmt.Println("here the levels")
	// To be done tomorrow !!!
	
	var dfsHelper func(current string, path []string)
	dfsHelper = func(current string, path []string) {
		if current == start {
			// Make a copy of the path and add it to trajectories
			// without this hadshi makaykhdmsh
			// ;)

			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			trajectories = append(trajectories, pathCopy)
			return
		}

		visited[current] = true
		for neighbor := range graph.Tunnels[current].Links {
			if !visited[neighbor] && levels[neighbor] < levels[current] {
				dfsHelper(neighbor, append([]string{neighbor}, path...))
			}
		}
		visited[current] = false
	}

	dfsHelper(graph.End_room.Name, []string{graph.End_room.Name})
	return trajectories
}

func SameLevel(node string, another_node string, level map[string]int) bool {
	return level[node] == level[another_node]
}

// Here there still some things to improve
// Finds the group of a specific path of a specific node
func FindOneGroup(graph *Colony, node string, shortest_path *Path) *Group {
	group := NewGroup()
	group.Shortest_Path = shortest_path
	group.Visited_Nodes[graph.Start_room.Name] = struct{}{}
	if shortest_path.Length != 0 {
		AddMapToAnotherMap(group.Visited_Nodes, shortest_path.Rooms_found[:len(shortest_path.Rooms_found)-1])
		group.Paths = append(group.Paths, shortest_path)
	}
	for _, key := range Priority(graph) {
		_, ok := group.Visited_Nodes[key]
		if key != node && !ok {
			path := NewPath()
			path.Rooms_found = group.BaseBFS(graph, key, graph.End_room.Name)
			path.Length = len(path.Rooms_found)
			if path.Length != 0 {
				AddMapToAnotherMap(group.Visited_Nodes, path.Rooms_found[:len(path.Rooms_found)-1])
				group.AppendPathToGroup(path)
			}

		}
	}

	return group
}

// Uses the func above to find all the groups if needed

// Function to find all the shortest Paths of the graph and store them in the colony
// don't know if it's a good thing to do for memory but just trying out

func FindAllGroups(graph *Colony) []*Group {
	groups := []*Group{}
	for _, node := range Priority(graph) {
		shortest_paths := DFS(graph, node)
		// hadi hna because not everytime ghanl9aw shortest b DFS khassna nrunniw BFS
		if len(shortest_paths) == 0 {
			shortest_paths = append(shortest_paths, BFS(graph, node))
		}

		for _, short := range shortest_paths {
			shortest_path := &Path{
				Rooms_found: short,
				Length:      len(short),
			}
			group := FindOneGroup(graph, node, shortest_path)
			groups = append(groups, group)

		}
	}

	return groups
}

// THE FUNCTION WHICH WILL find the best group step by step without leaking the memery

func FindTheBestGrp(graph *Colony) *Group {
	good_group := NewGroup()
	var is_first bool = false
	for _, node := range Priority(graph) {
		shortest_paths := DFS(graph, node)
		// hadi hna because not everytime ghanl9aw shortest b DFS khassna nrunniw BFS
		if len(shortest_paths) == 0 {
			shortest_paths = append(shortest_paths, BFS(graph, node))
		}

		for _, short := range shortest_paths {
			shortest_path := &Path{
				Rooms_found: short,
				Length:      len(short),
			}
			if !is_first {
				group := FindOneGroup(graph, node, shortest_path)
				group.CalculTurns(graph)
				good_group = group
				is_first = true
			} else {
				group := FindOneGroup(graph, node, shortest_path)
				group.CalculTurns(graph)
				good_group = Compare2Groups(graph, good_group, group)

			}

		}
	}

	return good_group
}

// This function is used to make us able to append paths in order
// so basically to not be needing again to sort them again
// when we access the group paths , they will be sorted from shortest to longest
func (G *Group) AppendPathToGroup(path_to_append *Path) {
	last := G.Paths[len(G.Paths)-1]
	if last.Length > path_to_append.Length {
		G.Paths = append(G.Paths[:len(G.Paths)-1], path_to_append)
		G.Paths = append(G.Paths, last)
	} else {
		G.Paths = append(G.Paths, path_to_append)
	}
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/(1024*1024))
}

//*********************************************************************************************************
// THIS SECTION FOR SOME FUNCTIONS THAT COULD BE USED LATER

// func (C *Colony) AppendPathToPATHS(path_to_append *Path) {
// 	last := C.Shortest_Paths[len(C.Shortest_Paths)-1]
// 	if last.Length > path_to_append.Length {
// 		C.Shortest_Paths = append(C.Shortest_Paths[:len(C.Shortest_Paths)-1], path_to_append)
// 		C.Shortest_Paths = append(C.Shortest_Paths, last)
// 	} else {
// 		C.Shortest_Paths = append(C.Shortest_Paths, path_to_append)
// 	}
// }

// func (C *Colony) FindAllShortestPATHS() {
// 	for child := range C.Tunnels[C.Start_room.Name].Links {
// 		shortest_paths := DFS(C, child)
// 		if len(shortest_paths) == 0 {
// 			shortest_paths = append(shortest_paths, BFS(C, child))
// 		}
// 		for _,path := range shortest_paths{
//            C.AppendPathToPATHS()
// 		}
// 	}
// }
