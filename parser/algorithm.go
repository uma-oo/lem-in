package parser




func BaseBFS(graph *Colony , start_node string , end_node string){


}


func ReconstructPath(graph *Colony, start_node string, end_node string){
	
}












func BfsShortestPath(graph *Colony, start_node string, end_node string) []string {
	parent_dict := make(map[string]string)
	current_layer := []string{}
	// next_layer := []string{}
    var current_node string
	current_layer = append(current_layer, start_node)
	parent_dict[start_node] = graph.Start_room.Name

	for len(current_layer) > 0 {
		current_node = Pop(current_layer)

		if current_node == end_node {
			return Reconstruct(parent_dict, start_node, end_node)
		}

	}

	return nil
}

func Pop(queue []string) string {
	popped := queue[0]
	queue = queue[1:]
	return popped
}

func Reconstruct(parent_map map[string]string, start_node string, end_node string) []string {
	path := []string{}
	current_node := end_node
	for current_node != "" {

	}
	return path
}

