package parser

import "fmt"

func DegreeNeighbors(graph *Colony) (string, map[string]struct{}) {
	var min string
	neigbors := make(map[string]struct{})
	var min_length int
	var initialized bool = false
	for element := range graph.Tunnels[graph.Start_room.Name].Links {
		fmt.Println(element)
		neigbors[element] = struct{}{}
		if !initialized {
			min_length = len(graph.Tunnels[element].Links)
			min = element
			initialized=true
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
