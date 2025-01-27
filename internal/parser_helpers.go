package internal

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

// tunnels checking :

func HandleTunnels(col *Colony, line_content []byte, line int) error {
	chunks, err := CheckTunnels(line, line_content)
	var ok bool
	if err != nil {
		return err
	} else {
		// here we check that the room alrzady exists in the rooms and ot
		_, ok1 := col.Rooms_coor[string(chunks[0])]
		_, ok2 := col.Rooms_coor[string(chunks[1])]
		ok = ok1 && ok2
		switch ok {
		case false:
			return errors.New("ERROR: invalid tunnel format, room does not exist at line: " + strconv.Itoa(line))
		case true:
			err := CheckConnections2(string(chunks[0]), string(chunks[1]), col)
			if err != nil {
				return errors.New("ERROR: invalid tunnel format, tunnel already exists at line: " + strconv.Itoa(line))
			}
		}
	}
	return nil
}

// had l function we check only that the tunnel has the correct format (mashi "" and not tied to itself )
func CheckTunnels(line int, line_content []byte) ([][]byte, error) {
	chunks := bytes.Split(line_content, []byte("-"))
	if len(chunks) != 2 {
		return nil, errors.New("ERROR: invalid data format, there is more than one connection at line: " + strconv.Itoa(line))
	} else if len(chunks) == 2 {
		if string(chunks[0]) == "" || string(chunks[1]) == "" {
			return nil, errors.New("ERROR: invalid data format, Invalid tunnel format at line:" + strconv.Itoa(line))
		} else if string(chunks[0]) == string(chunks[1]) {
			return nil, errors.New("ERROR: invalid data format, a connection tied to itself at line: " + strconv.Itoa(line))
		}
	}
	return chunks, nil
}

func CheckConnections2(room1_key string, room2_key string, col *Colony) error {
	value1, ok1 := col.Tunnels[room1_key]
	value2, ok2 := col.Tunnels[room2_key]
	if ok1 && ok2 {
		_, exists := value1.Links[value2.Name]
		if exists {
			// room2 deja kayna f links dyal room1 // error
			return errors.New("ERROR: invalid data format, this connection already exists")
		}
	} else if !(ok1 || ok2) {
		// hna bzuuj ma3mrhum tcreaw f links // awll mrra ybanu
		value1 = NewRoom()
		value2 = NewRoom()
		value1.setRoom(room1_key, col.Rooms_coor[room1_key][0], col.Rooms_coor[room1_key][1])
		value2.setRoom(room2_key, col.Rooms_coor[room2_key][0], col.Rooms_coor[room2_key][1])
	} else if !ok2 && ok1 {
		// room2 makaynash ga3 f tunnels
		value2 = NewRoom()
		value2.setRoom(room2_key, col.Rooms_coor[room2_key][0], col.Rooms_coor[room2_key][1])
	} else {
		// room1 bu7dha li makaynash
		value1 = NewRoom()
		value1.setRoom(room1_key, col.Rooms_coor[room1_key][0], col.Rooms_coor[room1_key][1])
	}
	value1.Links[room2_key] = struct{}{}
	value2.Links[room1_key] = struct{}{}
	col.Tunnels[room1_key] = value1
	col.Tunnels[room2_key] = value2
	return nil
}

// rooms handeling :
func CheckIsRoom(line_number int, line []byte) ([][]byte, bool) {
	chunks := bytes.Fields(line)
	if len(chunks) != 3 {
		return nil, false
	}
	return chunks, true
}

func isValidRoom(line_number int, chunks [][]byte) error {
	if !roomName.Match(chunks[0]) {
		return errors.New("ERROR: invalid data format, Invalid room Name at line: " + strconv.Itoa(line_number))
	} else if !roomCoordinates.Match(chunks[1]) || !roomCoordinates.Match(chunks[2]) {
		return errors.New("ERROR: invalid data format, Invalid room Coordinates: " + strconv.Itoa(line_number))
	} else if err := CheckCoorIsDuplicate(line_number, toInt(chunks[1]), toInt(chunks[2])); err != nil {
		return err
	}
	return nil
}

func toInt(bytes []byte) int {
	result := 0
	sign := 1
	for i, bt := range bytes {
		if string(bt) == "-" && i == 0 {
			sign = -1
			continue
		}
		result = result*10 + int(bt-'0')
	}
	return result * sign
}

func CheckCoorIsDuplicate(line int, x int, y int) error {
	value := szudzikPairSigned(x, y)
	if _, ok := combinaison[value]; ok {
		return errors.New("ERROR: invalid data format, The coordinates already exist at line " + strconv.Itoa(line))
	}
	combinaison[value] = struct{}{}
	return nil
}

func szudzikPair(x int, y int) int {
	if x >= y {
		return (x * x) + x + y
	}
	return (y * y) + x
}

// handles negative numbers and big sets too
func szudzikPairSigned(x int, y int) int {
	c := szudzikPair(convert(x), convert(y))
	if x < 0 || y < 0 {
		return -c - 1
	}
	return c
}

func convert(a int) int {
	if a >= 0 {
		return 2 * a
	}
	return (2 * a) - 1
}

// ants handeling

func CheckAnts(line_content []byte) bool {
	value, err := strconv.Atoi(string(line_content))
	return err == nil && value > 0
}

// comments and empty lines handeling:
func toIgnore(line_content []byte) bool {
	return CheckIsComment(line_content) || emptyLine.Match(line_content)
}

func CheckIsComment(line []byte) bool {
	return !start_command.Match(line) && !end_line.Match(line) && comment.Match(line)
}

// General :
func Error(mssg any) {
	fmt.Println("\033[31m" + fmt.Sprint(mssg) + "\033[0m")
}

func (c *Colony) CheckRooms() error {
	if c.Start_room.Name == "" && c.End_room.Name == "" {
		return errors.New("ERROR: invalid data format, start/end rooms not found")
	} else if c.Start_room.Name == "" {
		return errors.New("ERROR: invalid data format, start room not found")
	} else if c.End_room.Name == "" {
		return errors.New("ERROR: invalid data format, end room not found")
	}
	return nil
}

func (c *Colony) CheckStruct() error {
	if c.Ants == 0 {
		return errors.New("ERROR: invalid data format, No Ants found")
	} else if err := c.CheckRooms(); err != nil {
		return err
	} else if len(c.Rooms_coor) == 0 {
		return errors.New("ERROR: invalid data format, No rooms Found")
	} else if len(c.Tunnels) == 0 {
		return errors.New("ERROR: invalid data format, No Tunnels Found")
	}
	return nil
}
