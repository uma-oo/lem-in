package parser

import (
	"bytes"
	"errors"
	"strconv"
)

// fucntion to check if the the tunnel found fih ghir 2
func CheckTunnels(line int, line_content []byte) ([][]byte, error) {
	chunks := bytes.Split(line_content, []byte("-"))
	if len(chunks) != 2 {
		return nil, errors.New("ERROR: there is more or less than one connection at line: " + strconv.Itoa(line))
	} else if len(chunks) == 2 && string(chunks[0]) == string(chunks[1]) {
		return nil, errors.New("ERROR: a connection tied to itself at line: " + strconv.Itoa(line))
	}
	return chunks, nil
}

// function to check if the the pieces found rooms nit wlla walu

func HandleTunnels(col *Colony, line_content []byte, line int) error {
	chunks, err := CheckTunnels(line, line_content)
	var ok bool
	if err != nil {
		return err
	} else {
		_, ok1 := col.Rooms_coor[string(chunks[0])]
		_, ok2 := col.Rooms_coor[string(chunks[1])]

		ok = ok1 && ok2
		switch ok {
		case false:
			return errors.New("ERROR: the room in this tunnel doesn't exist at line: " + strconv.Itoa(line))
		case true:
			err := CheckConnections2(string(chunks[0]), string(chunks[1]), col)
			// err2 := CheckConnections2(string(chunks[1]), string(chunks[0]), col)
			if err != nil {
				return errors.New("ERROR: this connection already exists at line: " + strconv.Itoa(line))
			}
		}
	}
	return nil
}

// More cleaner function to check the links
// func to check if the rooms is already related to each other
func CheckConnections2(room1_key string, room2_key string, col *Colony) error {
	value1, ok1 := col.Tunnels[room1_key]
	value2, ok2 := col.Tunnels[room2_key]
	// check if room1 u room2 deja kaynin bzuuj
	// hna kayn 2 cases
	if ok1 && ok2 {
		_, exists := value1.Links[value2.Name]
		switch exists {
		case true:
			// room2 deja kayna f links dyal room1 // error
			return errors.New("ERROR: this connection already exists")
		case false:
			// room 2 makaynash -> nziiduha
			value1.Links[room2_key] = struct{}{}
			value2.Links[room1_key] = struct{}{}
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
