package parser

import (
	"errors"
	"fmt"
)

// after feeding the struct we need to make sure that we found and start and the end
// we cant' make sure til the EOF
// So after reading the file and be not able to reach an error catcheable while reading
// Empty the struct again

func (c *Colony) CheckStruct(cPt **Colony) error {
	if c.Start_room.Name == "" {
		*cPt = nil
		// this is working but why ???
		
		return errors.New("ERROR: The Start Command is Never Found")

	} else if c.End_room.Name == "" {
		*cPt = nil
		return errors.New("ERROR: The End Command is Never Found")
	} else if len(c.Rooms_coor) == 0 {
		*cPt = nil
		return errors.New("ERROR: No rooms Found")
	} else if len(c.Tunnels) == 0 {
		*cPt = nil
		return errors.New("ERROR: No Tunnels Found")
	}
	return nil
}

// Just to debug the colony
func (c *Colony) String() string {
	return fmt.Sprintf("Colony(Number of ants: %v, Start: %v, End: %v, Start Room: %v, End Room: %v , Rooms: %v )", c.Ants, c.Start, c.End, c.Start_room, c.End_room, c.Rooms_coor)
}
