package parser

import (
	"fmt"
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

func GetShortestPaths(graph *Colony, start string, target string) [][]string {
	paths := [][]string{}
	queue := [][]string{{start}}

	isVisisted := map[string]bool{}
	isVisisted[graph.Start_room.Name] = true
	isVisisted[start] = true

	for len(queue) > 0 {
		currentPath := queue[0]
		currentRoom := currentPath[len(currentPath)-1]
		queue = queue[1:]

		if currentRoom == graph.End_room.Name { // if found a path
			if len(paths) == 0 {
				paths = append(paths, currentPath)
				continue
			}
			if len(currentPath) == len(paths[0]) {
				paths = append(paths, currentPath)
				continue
			} else {
				break
			}
		}

		_, isLinkedToEnd := graph.Tunnels[currentRoom].Links[graph.End_room.Name]
		if isLinkedToEnd {
			newPath := append([]string{}, currentPath...)
			newPath = append(newPath, graph.End_room.Name)
			queue = append(queue, newPath)
			continue
		}

		for neighbor := range graph.Tunnels[currentRoom].Links {
			if !isVisisted[neighbor] {
				isVisisted[neighbor] = true
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	fmt.Println("shortest paths: ", paths)
	return paths
}

// func ReconstructPaths(graph *Colony, start string, target string) [][]string {
// 	paths_found := [][]string{}
// 	levels := Levels(graph, start, graph.End_room.Name)
// 	for neighbor := range graph.Tunnels[target].Links {
// 		if levels[neighbor] == levels[target]-1 {
// 			path := BFSOriented(graph, start, neighbor)
// 			fmt.Printf("path: %v\n", path)
// 			paths_found = append(paths_found, path)
// 		}
// 	}
// 	return paths_found
// }


// create a group of paths based on the shortest paths found
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
	for _, node := range Priority(graph) {
		// fmt.Println("Finding the paths for the node", node)
		// fmt.Println("Finding Groups")
		shortest_paths := GetShortestPaths(graph, node, graph.End_room.Name)
		// shortest_paths2 := GetShortestPaths()
		// hadi hna because not everytime ghanl9aw shortest b DFS khassna nrunniw BFS
		// if len(shortest_paths) == 0 {
		// 	shortest_paths = append(shortest_paths, BFS(graph, node))
		// }

		for _, short := range shortest_paths {
			shortest_path := &Path{
				Rooms_found: short,
				Length:      len(short),
			}
			// if we found a shortest path
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

// func printAlloc() {
// 	var m runtime.MemStats
// 	runtime.ReadMemStats(&m)
// 	fmt.Printf("%d MB\n", m.Alloc/(1024*1024))
// }

//*********************************************************************************************************

func BFSOriented(graph *Colony, start_node string, end_node string) []string {
	var current string
	trajectory := []string{}                    // The Path will be found
	var traversal *Traversal2 = NewTraversal2() // Initilaize dakshi lkula traversal
	traversal.isVisited[graph.Start_room.Name] = true
	traversal.Parent[start_node] = ""                        // element lwl visited
	start_element := SetNode(start_node)                     // kaykunu 3ndna string khasshum yt7wlu t structs li 3ndna
	traversal.Queue = append(traversal.Queue, start_element) // appendiw element lwlani l queue
	for len(traversal.Queue) > 0 {
		current = traversal.Pop()
		if current == graph.End_room.Name {
			for traversal.Parent[current] != "" { // base case ma7ddu mal9ash hadi donc mazal ma9ad l path
				trajectory = append([]string{current}, trajectory...)
				current = traversal.Parent[current]
			}
		}

		for element := range graph.Tunnels[current].Links {
			if _, ok := traversal.Parent[element]; !ok {
				if !traversal.isVisited[element] {
					traversal.isVisited[element] = true
					traversal.Parent[element] = current
					node_element := SetNode(element)
					traversal.Queue = append(traversal.Queue, node_element)
				}
			}
		}

	}
	if len(trajectory) != 0 || start_node == graph.End_room.Name { // edge case where the child of the start is the end itself or we didn't finc anything to add to the trajectory
		trajectory = append([]string{start_node}, trajectory...)
		return trajectory
	}

	return nil
}