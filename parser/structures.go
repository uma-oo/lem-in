package parser

import "regexp"

// what if we could create the type room like this

type Room struct {
	Name  string
	x     int
	y     int
	Links map[string]struct{} // this hack is from srm so useful !!!!! // string hna hya l key o hya room name dyal link
}

var combinaison = make(map[int]struct{})

type Colony struct {
	Ants       int
	Start      int
	End        int
	Start_room *Room
	End_room   *Room
	Rooms_coor map[string][]int
	Tunnels    map[string]*Room
}

// The Structs used in the process of finding the paths

type Node struct {
	Name *Room
}

type Traversal struct {
	Is_Visited   map[string]bool
	Visited_Node map[string][]string
	Queue        []*Node
}

type Group struct {
	Visited_Nodes map[string]struct{}
	Paths         []*Path
	Shortest_Path *Path

}

type Path struct {
	Rooms_found []string
	Length      int
}

type Traversal2 struct {
	Parent    map[string]string
	isVisited map[string]bool
	Queue     []*Node
}

var (
	start_line      = regexp.MustCompile(`^##start\s*$`)
	end_line        = regexp.MustCompile(`^##end\s*$`)
	comment         = regexp.MustCompile("^#")
	roomName        = regexp.MustCompile("^([^L#])[a-zA-Z0-9]*$")
	roomCoordinates = regexp.MustCompile("-?[0-9]+")
	emptyline       = regexp.MustCompile(`^\s*$`) // matches a empty line or line with spaces

)
