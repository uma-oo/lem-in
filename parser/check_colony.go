package parser

import (
	"errors"
	"fmt"
)

// after feeding the struct we need to make sure that we found and start and the end
// we cant' make sure til the EOF
// So after reading the file and be not able to reach an error catcheable while reading
// Empty the struct again

func (c *colony) CheckStruct(cPt **colony) error {
	if c.start_room.name == "" {
		*cPt = nil
		// this is working but why ???
		return errors.New("ERROR: The Start Command is Never Found")

	} else if c.end_room.name == "" {
		*cPt = nil
		return errors.New("ERROR: The End Command is Never Found")
	} else if len(c.rooms_coor) == 0 {
		*cPt = nil
		return errors.New("ERROR: No rooms Found")
	} else if len(c.Tunnels) == 0 {
		*cPt = nil
		return errors.New("ERROR: No Tunnels Found")
	}
	return nil
}

// Just to debug the colony
func (c *colony) String() string {
	return fmt.Sprintf("Colony(Number of ants: %v, Start: %v, End: %v, Start Room: %v, End Room: %v , Rooms: %v )", c.ants, c.start, c.end, c.start_room, c.end_room, c.rooms_coor)
}
