package parser

func BaseBFS(graph *Colony, start_node string, end_node string) ([]string, map[string]string) {
	var current string
	trajectory := []string{}                                 // The Path will be found
	var traversal *Traversal = NewTraversal()                // Initilaize dakshi lkula traversal
	traversal.Visited_Node[start_node] = ""                  // element lwl visited
	start_element := SetNode(start_node)                     // kaykunu 3ndna string khasshum yt7wlu t structs li 3ndna
	traversal.Queue = append(traversal.Queue, start_element) // appendiw element lwlani l queue

	for len(traversal.Queue) > 0 {
		current = traversal.Pop()
		if current == end_node {
			for traversal.Visited_Node[current] != "" {
				trajectory = append([]string{current}, trajectory...)
				current = traversal.Visited_Node[current]
			}
		}

		for element := range graph.Tunnels[current].Links {
			if _, ok := traversal.Visited_Node[element]; !ok {
				traversal.Visited_Node[element] = current
				node_element := SetNode(element)
				traversal.Queue = append(traversal.Queue, node_element)
			}
		}
	}
	return append([]string{start_node}, trajectory...), traversal.Visited_Node
}

func ReconstructPath(graph *Colony, start_node string, end_node string) {
}

func (T *Traversal) Pop() string {
	popped := T.Queue[0]
	T.Queue = T.Queue[1:]
	return popped.Name.Name
}
