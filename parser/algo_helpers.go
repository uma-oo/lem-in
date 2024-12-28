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

// wttf is this
// ana ktbt hadshi yes but how ????
// OMG sh7aal 7bbit

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

func AddMapToAnotherMap(whole map[string]interface{}, part map[string]string) {
	for key, _ := range part {
		whole[key] = struct{}{}
	}
}
