package parser

import "errors"

func (r *room) setRoom(name string, x int, y int) {
	r.name = name
	r.x = x
	r.y = y
}

// the function must check before adding a room to the colony

func (c *colony) addRoom(r ...*room) error {
	// check if the room will be adding exists already
	for _, ele := range r {
		if _, ok := c.rooms_coor[ele.name]; ok {
			// do something here
			return errors.New("ERROR: room is replicated")
		} else {
			c.rooms_coor[ele.name] = append(c.rooms_coor[ele.name], ele.x, ele.y)
		}
	}
	return nil
}
