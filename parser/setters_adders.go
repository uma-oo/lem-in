package parser

import "errors"

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
			// do something here
			return errors.New("ERROR: room is replicated")
		} else {
			c.Rooms_coor[ele.Name] = append(c.Rooms_coor[ele.Name], ele.x, ele.y)
		}
	}
	return nil
}
