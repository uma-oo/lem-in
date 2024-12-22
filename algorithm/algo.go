package algorithm

import (
	"fmt"
	p "lemin/parser"
)

// Let's implement DFS //with recursion :)
// Just one path at a time 
//hadshi s3iiiiiib

var marked = make(map[string]bool)
var Paths[][]string
var Path[]string
func FindPathsDFS(graph *p.Colony, start string, end string) [][]string {
	// end := graph.End_room.Name
	fmt.Println(marked)
	if val, ok := graph.Tunnels[start]; ok {
		marked[start]=true
		Path=append(Path, start)
		for key := range val.Links {
			if _, ok := marked[key]; !ok{
				fmt.Println("HERE", key)
				if key==end {
					Paths=append(Paths, Path)
					marked[start]=false
					marked[end]=false
					key=start
				}
				FindPathsDFS(graph, key, graph.End_room.Name)
			}
		}
		

	}
  return Paths
}

