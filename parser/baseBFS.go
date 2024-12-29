package parser

import "fmt"

func BaseBFS(graph *Colony, start_node string, end_node string) []string {
	var current string
	trajectory := []string{}                                 // The Path will be found
	var traversal *Traversal = NewTraversal()                // Initilaize dakshi lkula traversal
	traversal.Visited_Node[start_node] = ""                  // element lwl visited
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
			_, ok2 := Whole_traversal[element]
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
	return popped.Name.Name
}

func BFSOptimized(graph *Colony, start_node string, end_node string) [][]string {
	prioritized := Priority(graph)
	paths := [][]string{}
	Whole_traversal[start_node] = struct{}{}

	for i, element := range prioritized {
		fmt.Println(prioritized)
		path := BaseBFS(graph, element, end_node)
		if path != nil {
			paths = append(paths, path)
			AddMapToAnotherMap(Whole_traversal, path[:len(path)-1])
			i++
		} else if Contains(path, graph.Bad_Rooms[0]){
			DiscardPath(Whole_traversal, path)
			prioritized = append(prioritized, element)
			i=-1
		}

	}

	return paths
}
