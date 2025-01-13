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

	level[graph.Start_room.Name] = 0

	if start != graph.Start_room.Name {
		level[start] = 1
	}

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



func SameLevel(node string, another_node string, level map[string]int) bool {
	return level[node] == level[another_node]
}

func ReconstructPaths(graph *Colony, start string, target string) [][]string {
	paths_found := [][]string{}
	levels := Levels(graph, start, graph.End_room.Name)
	for neighbor := range graph.Tunnels[target].Links {
		if levels[neighbor] == levels[target]-1 {
			path := append(BFSOriented(graph, start, neighbor), target)
			paths_found = append(paths_found, path)
		}
	}
	return paths_found
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

// THE FUNCTION WHICH WILL find the best group step by step without leaking the memery

func FindTheBestGrp(graph *Colony) *Group {
	good_group := NewGroup()
	var is_first bool = false
	counter := 0
	for _, node := range Priority(graph) {
		fmt.Println("Finding the paths for the node", node)
		fmt.Println("Finding Groups")
		shortest_paths := ReconstructPaths(graph, node, graph.End_room.Name)
		// hadi hna because not everytime ghanl9aw shortest b DFS khassna nrunniw BFS
		// if len(shortest_paths) == 0 {
		// 	shortest_paths = append(shortest_paths, BFS(graph, node))
		// }
		fmt.Println("Found", counter)
		for _, short := range shortest_paths {
			shortest_path := &Path{
				Rooms_found: short,
				Length:      len(short),
			}
			// if we  found a shortest path
			counter++
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

func BFSOriented(graph *Colony, start_node string, end_node string) []string {
	var current string
	trajectory := []string{}                                 // The Path will be found
	var traversal *Traversal = NewTraversal()                // Initilaize dakshi lkula traversal
	traversal.Visited_Node[start_node] = []string{}          // element lwl visited
	start_element := SetNode(start_node)                     // kaykunu 3ndna string khasshum yt7wlu t structs li 3ndna
	traversal.Queue = append(traversal.Queue, start_element) // appendiw element lwlani l queue

	for len(traversal.Queue) > 0 {
		current = traversal.Pop()
		if current == end_node {
			for len(traversal.Visited_Node[current]) != 0 { // base case ma7ddu mal9ash hadi donc mazal ma9ad l path
				trajectory = append([]string{current}, trajectory...)
				for _, parent := range traversal.Visited_Node[current] {
					current = parent
				}

			}
		}

		for element := range graph.Tunnels[current].Links {

			_, ok1 := traversal.Visited_Node[element]

			if !ok1 {
				traversal.Visited_Node[element] = append(traversal.Visited_Node[element], current)
				node_element := SetNode(element)
				traversal.Queue = append(traversal.Queue, node_element)
			}
		}

	}
	if len(trajectory) != 0 || start_node == end_node { // edge case where the child of the start is the end itself or we didn't finc anything to add to the trajectory
		trajectory = append([]string{start_node}, trajectory...)
		return trajectory
	}

	return nil
}
//************************************************For Debuging******************************************************************************************//
func FindAllGroups(graph *Colony) []*Group {
	groups := []*Group{}
	for _, node := range Priority(graph) {
		shortest_paths := ReconstructPaths(graph, node, graph.End_room.Name)
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
			// printAlloc()
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