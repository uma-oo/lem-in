package internal

type Traversal struct {
	Visited_Node map[string]string
	Queue        []*Node
}

type Group struct {
	Visited_Nodes map[string]struct{}
	Paths         []*Path
	Shortest_Path *Path
	Turns         int
}

type Path struct {
	Rooms_found []string
	Length      int
}

type Node struct {
	*Room
}

// structs used for the printing process

type Agent struct {
	Pos      int
	PathUsed *Path
}

func NewTraversal() *Traversal {
	return &Traversal{
		Visited_Node: make(map[string]string),
		Queue:        []*Node{},
	}
}

func NewGroup() *Group {
	return &Group{
		Visited_Nodes: map[string]struct{}{},
	}
}

func NewPath() *Path {
	return &Path{
		Rooms_found: []string{},
	}
}

func NewAgent() *Agent {
	return &Agent{
		Pos:      0,
		PathUsed: NewPath(),
	}
}

func SetNode(name string) *Node {
	return &Node{
		Room: &Room{Name: name},
	}
}
