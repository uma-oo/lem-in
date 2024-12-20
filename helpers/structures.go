package helpers

import "regexp"

// what if we could create the type room like this

type room struct {
	name  string
	x     int
	y     int
	Links map[string]struct{} // this hack is from srm so useful !!!!! // string hna hya l key o hya room name dyal link
}

type coordinate struct {
	x int 
	y int 
}

var combinaison map[coordinate]struct{}



type colony struct {
	ants       int
	start      int
	end        int
	start_room *room
	end_room   *room
	rooms_coor map[string][]interface{}
	Tunnels    map[string]*room

	// tunnels    map[string][]string
}

var (
	start_line      = regexp.MustCompile(`^##start\s*$`)
	end_line        = regexp.MustCompile(`^##end\s*$`)
	comment         = regexp.MustCompile("^#")
	roomName        = regexp.MustCompile("^([^L#])[a-zA-Z0-9]*$")
	roomCoordinates = regexp.MustCompile("-?[0-9]+")
	emptyline       = regexp.MustCompile(`^\s*$`)   // matches a empty line or line with spaces 
	
)
