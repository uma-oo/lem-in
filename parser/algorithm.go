package parser

import "fmt"

// Let's implemnt the BFS without costs
// Just to test things out



var queue []string 

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

			if _, ok := visited[element]; !ok {
				visited[element] = current
			}
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

func BFSPathsForAnts(graph *Colony) []int{
	var current string
	var queue []string
	result := graph.Initialize(graph.Start_room.Name)
	visited := make(map[string]struct{})
	visited[graph.Start_room.Name] = struct{}{}
	proposed_explorations := []int{}
	frontier_map:=make(map[int][]string)
	for _, path := range result {
		
		frontier := path[len(path)-1]
		frontier_map[len(graph.Tunnels[frontier].Links)]=append(frontier_map[len(graph.Tunnels[frontier].Links)], frontier)
		proposed_explorations = append(proposed_explorations, len(graph.Tunnels[frontier].Links))
		
        
	}
	sorted_proposed_explorations:=quickSortStart(proposed_explorations)
	for _, element:= range sorted_proposed_explorations{
		queue = append(queue, frontier_map[element]...)
		current = Pop(queue)
		if _, ok := visited[current]; !ok {
            
		}
	}
	fmt.Println(queue)
    fmt.Println(frontier_map)
	return sorted_proposed_explorations
}

func (c *Colony) Initialize(key string) [][]string {
	var paths [][]string
	keys := ReturnKeys(c.Tunnels[key].Links)
	for i := 0; i < len(keys); i++ {
		paths = append(paths, []string{keys[i]})
	}
	return paths
}










// func Explore(graph *Colony, node string){
// 	for key := range graph.Tunnels[node].Links{
       
// 	}

// }





// here where will be choosing the correct node to be explored

	



// queue = append(queue, ReturnKeys(graph.Tunnels[graph.Start_room.Name].Links)...)
// 	for len(queue) > 0 {
// 		current = Pop(queue)
// 		result = append(result, current)
// 		visited[current] = graph.Start_room.Name
// 		queue = queue[1:]
// 		for _, element := range ReturnKeys(graph.Tunnels[current].Links) {

// 			// if !visited[element] && element != graph.End_room.Name {
// 			// 	visited[element] = true
// 			// 	queue = append(queue, element)
// 			// }

// 			if _ , ok := visited[element]; !ok {
//                 visited[element]=current

// 			}
// 		}
// 	}
// 	return result
