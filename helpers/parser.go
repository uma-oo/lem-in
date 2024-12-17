package helpers

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func NewColony() colony {
	return colony{
		start:      0,
		end:        0,
		start_room: NewRoom(),
		end_room:   NewRoom(),
		rooms_coor: make(map[string][]interface{}),
		tunnels:    make(map[string]*room),
	}
}

func NewRoom() *room {
	return &room{
		name:  "",
		x:     -1,
		y:     -1,
		Links: map[string]struct{}{},
	}
}

func (r *room) setRoom(name string, x int, y int) {
	r.name = name
	r.x = x
	r.y = y
}

func (r *room) setLinks(link map[string]struct{}) {
	r.Links = link
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

func Parse(filename string) (*colony, error) {
	// suppose checkinah
	// var error_parsing error
	colony := NewColony()
	var (
		start_found = false
		end_found   = false
	)

	file, err := os.Open(filename)
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	line := 0
	for scanner.Scan() {
		line++

		if line == 1 {
			if CheckAnts(scanner.Bytes()) {
				continue
			} else {
				return nil, errors.New("ERROR: invalid number of Ants")
			}
		} else {
			if CheckIsComment(line, scanner.Bytes()) {
				continue
			}
			if emptyline.Match(scanner.Bytes()) {
				continue
			} else if start_line.Match(scanner.Bytes()) && !start_found {
				start_found = true
				colony.start = line + 1
			} else if start_line.Match(scanner.Bytes()) && start_found {
				return nil, errors.New("ERROR: too many starts")
			} else if end_line.Match(scanner.Bytes()) && !end_found {
				end_found = true
				colony.end = line + 1
			} else if end_line.Match(scanner.Bytes()) && end_found {
				return nil, errors.New("ERROR: too many ends")
			} else if line == colony.start {
				ok, chunks := CheckIsRoom(line, scanner.Bytes())
				if ok {
					colony.start_room.setRoom(string(chunks[0]), toInt(chunks[1]), toInt(chunks[2]))
				} else {
					return nil, errors.New("ERROR: No start Found")
				}
			} else if line == colony.end {
				ok, chunks := CheckIsRoom(line, scanner.Bytes())
				if ok {
					colony.end_room.setRoom(string(chunks[0]), toInt(chunks[1]), toInt(chunks[2]))
					err := colony.addRoom(colony.start_room, colony.end_room)
					if err != nil {
						return nil, errors.New(string(scanner.Bytes()))
					}
				} else {
					return nil, errors.New("ERROR: No end Found")
				}
			} else if is_tunnel.Match(scanner.Bytes()) {
				err := HandleTunnels(&colony, scanner.Bytes(), line)
				if err != nil {
					return nil, err
				}
			} else {
				ok, chunks := CheckIsRoom(line, scanner.Bytes())
				if ok {
					new_room := NewRoom()
					new_room.setRoom(string(chunks[0]), toInt(chunks[1]), toInt(chunks[2]))
					err := colony.addRoom(new_room)
					if err != nil {
						return nil, errors.New(string(scanner.Bytes()))
					}
				}
			}

		}

	}
	return &colony, nil
}

func CheckAnts(line_content []byte) bool {
	_, err := strconv.Atoi(string(line_content))
	return err == nil
}

// check if the coordinates are also convertable
// checked so we are sure that they are digits

func CheckIsRoom(line_number int, line []byte) (bool, [][]byte) {
	if !CheckIsComment(line_number, line) {
		chunks := bytes.Split(line, []byte(" "))
		if len(chunks) > 3 {
			return false, nil
		}
		return roomName.Match(chunks[0]) && roomCoordinates.Match(chunks[1]) && roomCoordinates.Match(chunks[2]), chunks

	}
	return false, nil
}

func CheckIsComment(line_number int, line []byte) bool {
	return !start_line.Match(line) && !end_line.Match(line) && comment.Match(line)
}

func toInt(bytes []byte) int {
	result := 0
	for _, bt := range bytes {
		result = result*10 + int(bt-'0')
	}
	return result
}

func (c colony) String() string {
	return fmt.Sprintf("Colony(Start : %v, End: %v, Start Room: %v, End Room: %v , Rooms: %v , Links: %v)", c.start, c.end, c.start_room, c.end_room, c.rooms_coor, c.tunnels)
}

// slice is not optimized for this // done

// Check if 2 rooms have the same coordinates block of fcts
// use a hash function

// fucntion to check if the the tunnel found fih ghir 2
func CheckTunnels(line int, line_content []byte) ([][]byte, error) {
	chunks := bytes.Split(line_content, []byte("-"))
	if len(chunks) != 2 {
		return nil, errors.New("ERROR: there is more or less than one connection at line: " + strconv.Itoa(line))
	}
	return chunks, nil
}

// function to check if the the pieces found rooms nit wlla walu

func HandleTunnels(col *colony, line_content []byte, line int) error {
	chunks, err := CheckTunnels(line, line_content)
	fmt.Println(string(chunks[0]), string(chunks[1]))
	var ok bool
	if err != nil {
		return err
	} else {
		_, ok1 := col.rooms_coor[string(chunks[0])]
		_, ok2 := col.rooms_coor[string(chunks[1])]

		ok = ok1 && ok2

		switch ok {
		case false:
			return errors.New("ERROR: the room in this tunnel doesn't exist")
		case true:
			err := checkConnection(string(chunks[0]), string(chunks[1]), col)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// func to check if the rooms is already related to each other
// this is really getting more longer than expected
func checkConnection(room1_key string, room2_key string, col *colony) error {
	//
	if value, ok := col.tunnels[room1_key]; ok {
		// l9ina room2 deja kayna
		if _, ok := value.Links[room2_key]; ok {
			return errors.New("ERROR: this connection already exists")
		} else {
			fmt.Println("HERE")
			value.Links[room2_key] = struct{}{}
		}
	} else {
		value.Links[room2_key] = struct{}{}
		col.tunnels[room1_key] = value

		fmt.Println("THERE")
		// value.setLinks()
	}
	// we don't need to check if the room2 already contains room1 link because when adding, we add them both to each other
	return nil
}
