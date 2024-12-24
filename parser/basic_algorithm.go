package parser

// let's first initlize the paths

var solution *solver


func BFS(graph *Colony , solution *solver) *solver {
	current:=graph.Start_room.Name
	solution.queue=append(solution.queue, current)
	smallest := FindsSmallestDegree(current, graph)
	solution.paths=append(solution.paths,  []string{smallest})
	for len(solution.queue)>0{
		
	}

   return solution
}

// see the degree of each of the children of the current node we're exploring
// start with the smallest one and explore it
// add it to the queue

// function given some spefic parent finds the smallest degree node that should be explored

func FindsSmallestDegree(parent string, graph *Colony) string {
	var smallest string
	i := 0
	for link := range graph.Tunnels[parent].Links {
		if i == 0 {
			smallest = link
		} else if len(graph.Tunnels[link].Links) < len(graph.Tunnels[smallest].Links) {
			smallest = link
		}
		i++
	}
	return smallest
}



