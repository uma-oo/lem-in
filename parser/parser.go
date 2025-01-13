package parser

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) (*Colony, error) {
	// suppose checkinah
	// var error_parsing error
	colony := NewColony()
	var (
		start_found      = false
		end_found        = false
		start_room_found = false
		end_room_found   = false
	)

	file, err := os.Open(filename)
	scanner := bufio.NewScanner(file)
	if err != nil {
		return nil, errors.New("ERROR: Problem in the file")
	}
	defer file.Close()
	line := 0
	for scanner.Scan() {
		line++

		if line == 1 {
			if CheckAnts(scanner.Bytes()) {
				colony.Ants, _ = strconv.Atoi(string(scanner.Bytes()))
				continue
			} else {
				return nil, errors.New("ERROR: invalid number of Ants")
			}
		} else {
			if CheckIsComment(line, scanner.Bytes()) && start_found && !start_room_found {
				colony.Start = line + 1
			} else if CheckIsComment(line, scanner.Bytes()) && end_found && !end_room_found {
				colony.End = line + 1
			} else if CheckIsComment(line, scanner.Bytes()) {
				continue
			} else if emptyline.Match(scanner.Bytes()) && start_found && !start_room_found {
				colony.Start = line + 1
			} else if emptyline.Match(scanner.Bytes()) && end_found && !end_room_found {
				colony.End = line + 1
			} else if emptyline.Match(scanner.Bytes()) {
				continue
			} else if start_line.Match(scanner.Bytes()) && !start_found {
				start_found = true
				colony.Start = line + 1
			} else if start_line.Match(scanner.Bytes()) && start_found {
				return nil, errors.New("ERROR: too many starts at line: " + strconv.Itoa(line))
			} else if end_line.Match(scanner.Bytes()) && !end_found {
				end_found = true
				colony.End = line + 1
			} else if end_line.Match(scanner.Bytes()) && end_found {
				return nil, errors.New("ERROR: too many ends at line: " + strconv.Itoa(line))
			} else if !emptyline.Match(scanner.Bytes()) && !comment.Match(scanner.Bytes()) && line == colony.Start {
				start_room_found = true
				ok, chunks := CheckIsRoom(line, scanner.Bytes())
				if ok {
					colony.Start_room.setRoom(string(chunks[0]), toInt(chunks[1]), toInt(chunks[2]))
					err := colony.addRoom(colony.Start_room)
					if err != nil {
						return nil, err
					}
				} else {
					return nil, errors.New("ERROR: No start Found")
				}
			} else if !emptyline.Match(scanner.Bytes()) && !comment.Match(scanner.Bytes()) && line == colony.End {
				end_room_found = true
				ok, chunks := CheckIsRoom(line, scanner.Bytes())
				if ok {
					colony.End_room.setRoom(string(chunks[0]), toInt(chunks[1]), toInt(chunks[2]))
					err := colony.addRoom(colony.End_room)
					if err != nil {
						return nil, err
					}
				} else {
					return nil, errors.New("ERROR: No end Found")
				}

			} else if ok, chunks := CheckIsRoom(line, scanner.Bytes()); ok {
				err_coord := CheckCoorIsDuplicate(toInt(chunks[1]), toInt(chunks[2]))
				if err_coord != nil {
					return nil, err_coord
				}
				new_room := NewRoom()
				new_room.setRoom(string(chunks[0]), toInt(chunks[1]), toInt(chunks[2]))

				err := colony.addRoom(new_room)
				if err != nil {
					return nil, err
				}

			} else if strings.Contains(string(scanner.Bytes()), "-") {
				err := HandleTunnels(colony, scanner.Bytes(), line)
				if err != nil {
					return nil, err
				}
			}
		}

	}
	return colony, nil
}

// Check if 2 rooms have the same coordinates block of fcts
// use a hash function 
