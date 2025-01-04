package parser

func NewColony() *Colony {
	return &Colony{
		Ants:       0,
		Start:      0,
		End:        0,
		Start_room: NewRoom(),
		End_room:   NewRoom(),
		Rooms_coor: make(map[string][]int),
		Tunnels:    make(map[string]*Room),
	}
}

func NewRoom() *Room {
	return &Room{
		Links: map[string]struct{}{},
	}
}

func NewTraversal() *Traversal {
	return &Traversal{
		Visited_Node: make(map[string]string),
		Queue:        []*Node{},
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
