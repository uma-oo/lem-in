package parser

func (G *Group) BaseBFS(graph *Colony, start_node string, end_node string) []string {
	var current string
	trajectory := []string{}                  // The Path will be found
	var traversal *Traversal = NewTraversal() // Initilaize dakshi lkula traversal
	G.Visited_Nodes[graph.Start_room.Name] = struct{}{}
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
			_, ok2 := G.Visited_Nodes[element]
			if !ok1 && !ok2 {
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

func (T *Traversal) Pop() string {
	popped := T.Queue[0]
	T.Queue = T.Queue[1:]
	return popped.Name.Name
}

func (T *Traversal2) Pop() string {
	popped := T.Queue[0]
	T.Queue = T.Queue[1:]
	return popped.Name.Name
}


// Doesn't belong to the Group BY ANY WAY
func BFS(graph *Colony, start_node string) []string {
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


