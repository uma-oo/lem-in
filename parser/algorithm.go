package parser

import "fmt"

// func BFS(graph *Colony) {
// 	var Paths [][]string
// 	visited := make(map[string]struct{})
// 	start_points := graph.Tunnels[graph.Start_room.Name].Links
// 	visited[graph.Start_room.Name] = struct{}{}
// 	for key := range start_points {
// 		slice := []string{}
// 		current := key
// 		slice = append(slice, current)
// 		queue := []string{}
// 		for link := range graph.Tunnels[current].Links {
// 			queue = append(queue, link)
// 			if _, ok := visited[link]; !ok && key != graph.End_room.Name {
// 				slice = append(slice, link)
// 				visited[link] = struct{}{}
// 			}
// 		}
// 		fmt.Println("slice", slice)
// 		Paths = append(Paths, slice)
// 	}
// 	fmt.Println(start_points)
// 	fmt.Println(Paths)
// }

func BFSOnSingleNode(graph *Colony, node string) {
	var paths [][]string
	queue := []struct {
		node string
		path []string
	}{{
		node: node,
		path: []string{node},
	}}
	visited := make(map[string]struct{})
	visited[node] = struct{}{} // Mark the starting node as visited

	for len(queue) > 0 {
		// Dequeue the front node
		current := queue[0]
		queue = queue[1:]

		// Process the current node
		for link := range graph.Tunnels[current.node].Links {
			if _, ok := visited[link]; !ok {
				// Mark the link as visited
				visited[link] = struct{}{}

				// Create a new path that includes the current node and the link
				newPath := append([]string{}, current.path...) 
				newPath = append(newPath, link)           

				// Check if the link is the end node
				if link == graph.End_room.Name {
					// If it's the end node, add the path to the paths slice
					paths = append(paths, newPath)
				} else {
					// Otherwise, enqueue the new node and the updated path
					queue = append(queue, struct {
						node string
						path []string
					}{node: link, path: newPath})
				}
			}
		}
	}

	fmt.Println("Visited nodes:", visited)
	fmt.Println("Paths found:", paths)
}
