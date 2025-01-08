package parser

func NewColony() *Colony {
	return &Colony{
		Ants:           0,
		Start:          0,
		End:            0,
		Start_room:     NewRoom(),
		End_room:       NewRoom(),
		Rooms_coor:     make(map[string][]int),
		Tunnels:        make(map[string]*Room),
		Shortest_Paths: []*Path{},
	}
}

func NewRoom() *Room {
	return &Room{
		Links: map[string]struct{}{},
	}
}

func NewTraversal() *Traversal {
	return &Traversal{
		Visited_Node: make(map[string][]string),
		Is_Visited:   make(map[string]bool),
		Queue:        []*Node{},
	}
}

func NewTraversal2() *Traversal2 {
	return &Traversal2{
		Parent:    make(map[string]string),
		isVisited: make(map[string]bool),
		Queue:     []*Node{},
	}
}

func NewPath() *Path {
	return &Path{
		Rooms_found: []string{},
	}
}

func NewGroup() *Group {
	return &Group{
		Visited_Nodes: map[string]struct{}{},
	}
}

func NewAgent() *Agent {
	return &Agent{
		Name:       0,
		PathUsed:   NewPath(),
		HasArrived: false,
		Start_Step: 0,
		End_Step: 0,
	}
}
