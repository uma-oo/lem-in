package parser

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

func Error(mssg any) {
	fmt.Println("\033[31m" + fmt.Sprint(mssg) + "\033[0m")
}

func CheckAnts(line_content []byte) bool {
	value, err := strconv.Atoi(string(line_content))
	return err == nil && value > 0
}

// check if the coordinates are also convertable
// checked so we are sure that they are digits

func CheckIsRoom(line_number int, line []byte) (bool, [][]byte) {
	if !CheckIsComment(line_number, line) {
		chunks := bytes.Fields(line)
		if len(chunks) > 3 {
			return false, nil
		}
		return roomName.Match(chunks[0]) && roomCoordinates.Match(chunks[1]) && roomCoordinates.Match(chunks[2]), chunks

	}
	return false, nil
}

func CheckCoorIsDuplicate(x int, y int) error {
	value := szudzikPairSigned(x, y)
	if _, ok := combinaison[value]; ok {
		return errors.New("ERROR: The coordinates of this room already exist")
	}
	combinaison[value] = struct{}{}
	return nil
}

func CheckIsComment(line_number int, line []byte) bool {
	return !start_line.Match(line) && !end_line.Match(line) && comment.Match(line)
}

// need to add negative numbers  -> added :)

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

func (c *Colony) Degree(room string) int {
	return len(c.Tunnels[room].Links) - 1
}

func Contains(path []string, room string) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}

// Just to debug the colony
func (c *Colony) String() string {
	return fmt.Sprintf("Colony(Number of ants: %v, Start: %v, End: %v, Start Room: %v, End Room: %v , Tunnels:%v)", c.Ants, c.Start, c.End, c.Start_room, c.End_room, c.Tunnels)
}

func (r *Room) String() string {
	return fmt.Sprintf("Room %v and its links are %v", r.Name, r.Links)
}

// To debug the agent again and again
func (a *Agent) String() string {
	return fmt.Sprintf("Ant Position in the path: %v, Path used: %v", a.Pos, a.PathUsed.String())
}

func (t *Traversal) String() string {
	return fmt.Sprintf("Is_visited map: %v, Visited_nodes: %v queue: %v", t.Is_Visited, t.Visited_Node, t.Queue)
}

// this function is here to just show us the links :<)=

func (c *Colony) PrintLinks(links map[string]*Room) {
	for key, value := range links {
		fmt.Printf("the room  %s and the links are %v \n", key, value.Links)
	}
}

func (P *Path) String() []string {
	return P.Rooms_found
}

func (G *Group) String() [][]string {
	paths := [][]string{}
	for _, path := range G.Paths {
		paths = append(paths, path.Rooms_found)
	}
	return paths
}
