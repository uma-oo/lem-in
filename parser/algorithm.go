package parser

// Let's implemnt the BFS without costs
// Just to test things out

func BFSPaths(graph *Colony) []string {
	var current string
	var queue []string
	var result []string
	visited := make(map[string]string)
	queue = append(queue, ReturnKeys(graph.Tunnels[graph.Start_room.Name].Links)...)
	for len(queue) > 0 {
		current = Pop(queue)
		result = append(result, current)
		visited[current] = graph.Start_room.Name
		queue = queue[1:]
		for _, element := range ReturnKeys(graph.Tunnels[current].Links) {

			// if !visited[element] && element != graph.End_room.Name {
			// 	visited[element] = true
			// 	queue = append(queue, element)
			// }
		}
	}
	return result
}

func ReturnKeys(m map[string]struct{}) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func Pop(queue []string) string {
	return string(queue[0])
}

// func BFSPathsForAnts(graph *p.Colony){

// }

func (c *Colony) Initialize() [][]string {
	var paths [][]string
	for i, j := 0, 0; i < len(ReturnKeys(c.Tunnels[c.Start_room.Name].Links)); i, j = i+1, j+1 {
		paths = append(paths, []string{ReturnKeys(c.Tunnels[c.Start_room.Name].Links)[i]})
	}
	return paths
}


