package parser

import (
	"errors"
	"fmt"
)

// after feeding the struct we need to make sure that we found and start and the end
// we cant' make sure til the EOF
// So after reading the file and be not able to reach an error catcheable while reading
// Empty the struct again

func (c *Colony) CheckStruct() error {
	fmt.Println(c.String())
	if c.Start_room.Name == "" {
		return errors.New("ERROR: The Start Command is Never Found")
	} else if c.End_room.Name == "" {
		return errors.New("ERROR: The End Command is Never Found")
	} else if len(c.Rooms_coor) == 0 {
		return errors.New("ERROR: No rooms Found")
	} else if len(c.Tunnels) == 0 {
		return errors.New("ERROR: No Tunnels Found")
	} else if len(c.Tunnels[c.Start_room.Name].Links) == 0 {
		return errors.New("ERROR: The Start Room is not tied")
	} else if len(c.Tunnels[c.End_room.Name].Links) == 0 {
		return errors.New("ERROR: The End Room is not tied")
	}
	return nil
}
