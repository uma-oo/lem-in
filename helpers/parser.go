package helpers

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// what if we could create the type room like this

type room struct {
	name string
	x int
	y int 

}

type colony struct {
	start      int
	end        int
	start_room room
	end_room   room
	rooms      map[string][]interface{}
	// tunnels []tunnel
}

var (
	start_line      = regexp.MustCompile("^##start$")
	end_line        = regexp.MustCompile("^##end$")
	comment         = regexp.MustCompile("^#")
	roomName        = regexp.MustCompile("^([^L#])[a-zA-Z0-9]*$")
	roomCoordinates = regexp.MustCompile("[0-9]+")
)

func NewColony() colony {
	return colony{
		start:      0,
		end:        0,
		start_room: NewRoom(),
		end_room:   NewRoom(),
		rooms:      make(map[string][]interface{}, 0),
	}
}

func NewRoom() room {
	return room{
		name :"",
		x:-1,
		y:-1,
	}
}

// the function must check before adding a room to the colony

func (c *colony) addRoom(r room) error {
	for key  := range c.rooms {
		if _, ok := c.rooms[key]; ok {
			return errors.New("Error: there is a replicate room")
		} 

	}
	c.rooms[room.name]=[]int{room.x, room.y}
	return nil 
}

func (r *room) setRoom(key string, value ...interface{}) {
	r.room[key] = value
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
			if start_line.Match(scanner.Bytes()) && !start_found {
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
					colony.addRoom(colony.start_room, colony.end_room)
				} else {
					return nil, errors.New("ERROR: No end Found")
				}
			} else {
				ok, chunks := CheckIsRoom(line, scanner.Bytes())
				if ok {
					new_room := NewRoom()
					new_room.setRoom(string(chunks[0]), toInt(chunks[1]), toInt(chunks[2]))
					colony.rooms = append(colony.rooms, new_room)
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
	return fmt.Sprintf("Colony(Start : %v, End: %v, Start Room: %v, End Room: %v , Rooms: %v)", c.start, c.end, c.start_room, c.end_room, c.rooms)
}

// slice is not optimized for this
