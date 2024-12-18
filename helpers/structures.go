package helpers

import "regexp"

// what if we could create the type room like this

type room struct {
	name  string
	x     int
	y     int
	Links map[string]struct{} // this hack is from srm so useful !!!!! // string hna hya l key o hya room name dyal link
}

type colony struct {
	start      int
	end        int
	start_room *room
	end_room   *room
	rooms_coor map[string][]interface{}
	tunnels    map[string]*room

	// tunnels    map[string][]string
}

var (
	start_line      = regexp.MustCompile("^##start$")
	end_line        = regexp.MustCompile("^##end$")
	comment         = regexp.MustCompile("^#")
	roomName        = regexp.MustCompile("^([^L#])[a-zA-Z0-9]*$")
	roomCoordinates = regexp.MustCompile("[0-9]+")
	emptyline       = regexp.MustCompile(`^\\s*$`)
	// is_tunnel       = regexp.MustCompile(`^([a-zA-Z0-9]+)[-]([a-zA-Z0-9]+)$`)
)
