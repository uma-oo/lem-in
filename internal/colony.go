package internal

import (
	"errors"
	"regexp"
)

type Room struct {
	Name  string
	x, y  int                 // Room coordinates
	Links map[string]struct{} // Room links to other rooms
}

var combinaison = make(map[int]struct{})

type Colony struct {
	Ants           int
	Start          int
	End            int
	Start_room     *Room
	End_room       *Room
	Rooms_coor     map[string][]int
	Tunnels        map[string]*Room
}

var (
	start_command   = regexp.MustCompile(`^##start$`)
	end_line        = regexp.MustCompile(`^##end$`)
	comment         = regexp.MustCompile(`^#`)
	tunnel          = regexp.MustCompile(`^[^L\-#]+-[^L\-#]+$`)
	roomName        = regexp.MustCompile(`^[^L#]*$`)
	roomCoordinates = regexp.MustCompile(`^-?[0-9]+$`)
	emptyLine       = regexp.MustCompile(`^$`)
)

func NewColony() *Colony {
	return &Colony{
		Ants:           0,
		Start:          0,
		End:            0,
		Start_room:     NewRoom(),
		End_room:       NewRoom(),
		Rooms_coor:     make(map[string][]int),
		Tunnels:        make(map[string]*Room),
	}
}

func NewRoom() *Room {
	return &Room{
		Name:  "",
		Links: map[string]struct{}{},
	}
}

func (r *Room) setRoom(name string, x int, y int) {
	r.Name = name
	r.x = x
	r.y = y
}

// the function must check before adding a room to the colony
func (c *Colony) addRoom(r ...*Room) error {
	// check if the room will be adding exists already
	for _, ele := range r {
		if _, ok := c.Rooms_coor[ele.Name]; ok {
			return errors.New("ERROR: invalid data format, room is replicated")
		} else {
			c.Rooms_coor[ele.Name] = append(c.Rooms_coor[ele.Name], ele.x, ele.y)
		}
	}
	return nil
}
