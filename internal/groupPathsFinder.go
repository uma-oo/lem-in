package internal

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
	return paths
}

func BuildNewGroup(graph *Colony, node string, shortest_path *Path) *Group {
	group := NewGroup()
	group.Shortest_Path = shortest_path
	group.Visited_Nodes[graph.Start_room.Name] = struct{}{}
	group.Paths = append(group.Paths, shortest_path)
	MarkRoomsVisited(group.Visited_Nodes, shortest_path.Rooms_found[:len(shortest_path.Rooms_found)-1])

	for key := range graph.Tunnels[graph.Start_room.Name].Links {
		_, ok := group.Visited_Nodes[key]
		if key != node && !ok {
			path := NewPath()
			path.Rooms_found = group.BaseBFS(graph, key, graph.End_room.Name)
			path.Length = len(path.Rooms_found)
			if path.Length != 0 {
				MarkRoomsVisited(group.Visited_Nodes, path.Rooms_found[:len(path.Rooms_found)-1])
				group.AppendPathToGroup(path)
			}
		}
	}
	return group
}

func FindTheBestGrp(graph *Colony) *Group {
	good_group := NewGroup()
	is_first := true
	for node := range graph.Tunnels[graph.Start_room.Name].Links {
		shortest_paths := GetShortestPaths(graph, node, graph.End_room.Name)
		for _, short := range shortest_paths {
			shortest_path := &Path{
				Rooms_found: short,
				Length:      len(short),
			}
			group := BuildNewGroup(graph, node, shortest_path)
			group.CalculTurns(graph)
			if is_first {
				good_group = group
				is_first = false
			} else {
				good_group = Compare2Groups(graph, good_group, group)
			}
		}
	}
	return good_group
}

func (G *Group) AppendPathToGroup(path_to_append *Path) {
	last := G.Paths[len(G.Paths)-1]
	if last.Length > path_to_append.Length {
		G.Paths = append(G.Paths[:len(G.Paths)-1], path_to_append)
		G.Paths = append(G.Paths, last)
	} else {
		G.Paths = append(G.Paths, path_to_append)
	}
}

func (G *Group) BaseBFS(graph *Colony, start_node string, end_node string) []string {
	var current string
	trajectory := []string{}    // The Path will be found
	traversal := NewTraversal() // Initilaize dakshi lkula traversal
	G.Visited_Nodes[graph.Start_room.Name] = struct{}{}
	traversal.Visited_Node[start_node] = ""        // element lwl visited
	start_element := SetNode(start_node)                     // kaykunu 3ndna string khasshum yt7wlu t structs li 3ndna
	traversal.Queue = append(traversal.Queue, start_element) // appendiw element lwlani l queue

	for len(traversal.Queue) > 0 {
		current = traversal.Pop()
		if current == end_node {
			for traversal.Visited_Node[current] != "" { // base case ma7ddu mal9ash hadi donc mazal ma9ad l path
				trajectory = append([]string{current}, trajectory...)
					current = traversal.Visited_Node[current]
			}
		}
		for element := range graph.Tunnels[current].Links {
			_, ok1 := traversal.Visited_Node[element]
			_, ok2 := G.Visited_Nodes[element]
			if !ok1 && !ok2 {
				traversal.Visited_Node[element] = current
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

func (T *Traversal) Pop() string {
	popped := T.Queue[0]
	T.Queue = T.Queue[1:]
	return popped.Name
}
