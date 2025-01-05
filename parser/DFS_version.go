package parser

import "fmt"

// used BFS hna to determine levels
// this is the implementation based on this stackoverflow question
// link hahuwa for more details : https://stackoverflow.com/questions/14144071/finding-all-the-shortest-paths-between-two-nodes-in-unweighted-undirected-graph?rq=1
// BFS + reverse DFS
func Levels(graph *Colony, start string, end string) map[string]int {
	level := make(map[string]int)
	var current string
	var traversal *Traversal = NewTraversal()                // Initilaize dakshi lkula traversal
	traversal.Visited_Node[start] = []string{}               // element lwl visited
	start_element := SetNode(start)                          // kaykunu 3ndna string khasshum yt7wlu t structs li 3ndna
	traversal.Queue = append(traversal.Queue, start_element) // appendiw element lwlani l queue
	level[start] = 1
	level[graph.Start_room.Name]=0

	for len(traversal.Queue) > 0 {
		current = traversal.Pop()
		if current == end {
			break
		}
    
		for element := range graph.Tunnels[current].Links {
			if _, ok := traversal.Visited_Node[element]; !ok {
				traversal.Visited_Node[element] = append(traversal.Visited_Node[element], current)
				level[element] = level[current] + 1
				node_element := SetNode(element)
				traversal.Queue = append(traversal.Queue, node_element)
				level[element] = level[current] + 1
			} else {
				for _, parent := range traversal.Visited_Node[element] {
					if SameLevel(parent, current, level) {
						traversal.Visited_Node[element] = append(traversal.Visited_Node[element], current)
					}
				}
			}
		}
	}

	return level
}

// Backtracking bash n3awdu njbduu all the paths based on the levels we already have
// we can have more than one shortest path
func DFS(graph *Colony, start string) [][]string {
	var trajectories [][]string
	visited := make(map[string]bool)
	visited[graph.Start_room.Name]=true
	levels := Levels(graph, start, graph.End_room.Name)
	fmt.Printf("levels: %v\n", levels)
	var dfsHelper func(current string, path []string)
	dfsHelper = func(current string, path []string) {
		if current == start {
			// Make a copy of the path and add it to trajectories
			// without this hadshi makaykhdmsh 
			// ;) 
			fmt.Printf("\"here\": %v\n", "here")
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			trajectories = append(trajectories, pathCopy)
			return
		}
        fmt.Printf("\"There\": %v\n", "There")
		fmt.Printf("current: %v\n", current)
		visited[current] = true
		for neighbor := range graph.Tunnels[current].Links {
			fmt.Printf("neighbor: %v\n", neighbor)
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


// Here where we find all the groups that could be used in the solution 

func FindAllGroups(graph *Colony) []*Group {
	groups := []*Group{}
	for _, node := range Priority(graph) {
		shortest_paths := DFS(graph, node)
		if len(shortest_paths)==0{
			shortest_paths = append(shortest_paths, BFS(graph, node))
		}
		for _, short := range shortest_paths {
			group := NewGroup()
			group.Visited_Nodes[graph.Start_room.Name]=struct{}{}
			path := NewPath()
			path.Rooms_found = short
			path.Length = len(path.Rooms_found)
			if path.Length != 0 {
				AddMapToAnotherMap(group.Visited_Nodes, path.Rooms_found[:len(path.Rooms_found)-1])
				group.Paths = append(group.Paths, path)
			}
			for _, key := range Priority(graph) {
				if key != node {
					path = NewPath()
					path.Rooms_found = group.BaseBFS(graph, key, graph.End_room.Name)
					path.Length = len(path.Rooms_found)
					if path.Length != 0 {
						AddMapToAnotherMap(group.Visited_Nodes, path.Rooms_found[:len(path.Rooms_found)-1])
						group.Paths = append(group.Paths, path)
					}

				}
			}
			groups = append(groups, group)

		}
	}

	return groups
}
