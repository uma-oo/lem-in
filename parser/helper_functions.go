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
		chunks := bytes.Split(line, []byte(" "))
		if len(chunks) > 3 {
			return false, nil
		}
		return roomName.Match(chunks[0]) && roomCoordinates.Match(chunks[1]) && roomCoordinates.Match(chunks[2]), chunks

	}
	return false, nil
}

func CheckCoorIsDuplicate(x int, y int) error {
	value:=szudzikPairSigned(x, y)
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

func Contains(path []string , room string)  bool {
	for _, r := range path {
		if r==room{
			return true 
		}
	}
	return false 
}