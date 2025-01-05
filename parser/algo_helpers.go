package parser

import (
	"math"
)

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

// wttf is this
// ana ktbt hadshi yes but how ????
// OMG sh7aal 7bbit
// Modify this to include the Bad Room too
// if a path found contains the bad ROOM
// we run the BFS on the OTHER One
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

func DiscardPath(whole map[string]interface{}, path []string) {
	for _, room := range path {
		delete(whole, room)
	}
}

// the bad room is always the first one to accur and second one and so on

func DetectBadRooms(graph *Colony) {
	initiliazed := false
	var max int
	for room, length := range graph.Tunnels {
		if room == graph.End_room.Name || room == graph.Start_room.Name {
			continue
		} else {
			if len(length.Links) > int(AverageRoomLinks(graph)) && !initiliazed {
				graph.Bad_Rooms = append(graph.Bad_Rooms, room)
				max = len(length.Links)
				initiliazed = true

			} else if len(length.Links) > int(AverageRoomLinks(graph)) && initiliazed {
				if len(length.Links) > max {

					max = len(length.Links)
					graph.Bad_Rooms = append([]string{room}, graph.Bad_Rooms...)
				} else {
					graph.Bad_Rooms = append(graph.Bad_Rooms, room)
				}
			}
		}
	}
}

func AverageRoomLinks(graph *Colony) float64 {
	var average int
	for _, properties := range graph.Tunnels {
		average += len(properties.Links)
	}
	return math.Round(float64(average / (len(graph.Tunnels)))) // -2 here is to exclude the start and the end
}

// func PriorityWithBadRoom(graph *Colony) []string {
// 	var Path []string
// 	new_arr_priority := []string{}
// 	for _, element := range Priority(graph) {
// 		fmt.Println("element", element)
// 		Path = BFS(graph, element, graph.End_room.Name)
// 		fmt.Println("Path Found", Path, "element", element)
// 		if Contains(Path, graph.Bad_Rooms[0]) {
// 			continue
// 		} else {
// 			new_arr_priority = append(new_arr_priority, element)
// 		}
// 	}
// 	return new_arr_priority
// }

func (P *Path) String() []string {
	return P.Rooms_found
}

func (G *Group) String() [][]string {
	paths := [][]string{}
	for _, path := range G.Paths {
		paths = append(paths, path.Rooms_found)
	}
	return paths
}
