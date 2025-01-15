package parser

func DegreeNeighbors(graph *Colony) (string, map[string]struct{}) {
	var min string
	neigbors := make(map[string]struct{})
	var min_length int
	var initialized bool = false
	for element := range graph.Tunnels[graph.Start_room.Name].Links {
		neigbors[element] = struct{}{}
		if !initialized {
			min_length = len(graph.Tunnels[element].Links)
			min = element
			initialized = true
		} else {
			if min_length >= len(graph.Tunnels[element].Links) {
				min_length = len(graph.Tunnels[element].Links)
				min = element
			}
		}
	}
	delete(neigbors, min)
	return min, neigbors
}

// This function forget why it's here !!!!!
// Eeeh we need to apply it multiple times so the first one we don't have anything
// The second attempt we already seen the first one and we need to see again to which room the exploration is prioritized

func DegreeNeighborsTwo(map_priority map[string]struct{}, graph *Colony) (string, map[string]struct{}) {
	var min string
	neigbors := make(map[string]struct{})
	var min_length int
	var initialized bool = false
	for element := range map_priority {
		neigbors[element] = struct{}{}
		if !initialized {
			min_length = len(graph.Tunnels[element].Links)
			min = element
			initialized = true
		} else {
			if min_length >= len(graph.Tunnels[element].Links) {
				min_length = len(graph.Tunnels[element].Links)
				min = element
			}
		}
	}

	delete(neigbors, min)
	return min, neigbors
}

func Priority(graph *Colony) []string {
	arr_priority := []string{}
	var (
		min          string
		map_priority = make(map[string]struct{})
	)
	for i := 0; i < len(graph.Tunnels[graph.Start_room.Name].Links); i++ {
		if i == 0 {
			min, map_priority = DegreeNeighbors(graph)
			arr_priority = append(arr_priority, min)
		} else {
			min, map_priority = DegreeNeighborsTwo(map_priority, graph)
			arr_priority = append(arr_priority, min)
		}
	}
	return arr_priority
}

func AddMapToAnotherMap(whole map[string]struct{}, part []string) {
	for _, element := range part {
		whole[element] = struct{}{}
	}
}
