package internal

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) (*Colony, error) {
	colony := NewColony()
	var (
		ants_found          = false
		start_command_found = false
		end_command_found   = false
		start_room_found    = false
		end_room_found      = false
		tunnel_found        = false
	)
	
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("ERROR: file does not exists")
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()
	line := 0
	for scanner.Scan() {
		str := strings.TrimSpace(scanner.Text())
		text := []byte(str)
		line++
		if !ants_found {
			if toIgnore(text) {
				continue
			}
			if CheckAnts(text) {
				colony.Ants, _ = strconv.Atoi(string(text))
				ants_found = true
			} else {
				return nil, errors.New("ERROR: invalid data format, invalid number of Ants")
			}
		} else {
			if toIgnore(text) && start_command_found && !start_room_found {
				colony.Start = line + 1
			} else if toIgnore(text) && end_command_found && !end_room_found {
				colony.End = line + 1
			} else if toIgnore(text) {
				continue
			} else if start_command.Match(text) && !start_command_found {
				start_command_found = true
				colony.Start = line + 1
			} else if start_command.Match(text) && start_command_found {
				return nil, errors.New("ERROR: invalid data format, too many start commands at line: " + strconv.Itoa(line))
			} else if end_line.Match(text) && !end_command_found {
				end_command_found = true
				colony.End = line + 1
			} else if end_line.Match(text) && end_command_found {
				return nil, errors.New("ERROR: invalid data format, too many end commands at line: " + strconv.Itoa(line))
			} else if !toIgnore(text) && line == colony.Start && !end_command_found {
				if colony.Start_room.Name != "" {
					return nil, errors.New("ERROR: invalid data format, Can't have multiple start rooms")
				} else {
					chunks, ok := CheckIsRoom(line, text)
					if ok {
						err := isValidRoom(line, chunks)
						if err != nil {
							return nil, err
						}
						colony.Start_room.setRoom(string(chunks[0]), toInt(chunks[1]), toInt(chunks[2]))
						err = colony.addRoom(colony.Start_room)
						if err != nil {
							return nil, err
						}
						start_room_found = true
					} else {
						return nil, errors.New("ERROR: invalid data format, Invalid room format at line: " + strconv.Itoa(line))
					}
					start_command_found = false
				}
			} else if !toIgnore(text) && line == colony.End && !start_command_found {
				if colony.End_room.Name != "" {
					return nil, errors.New("ERROR: invalid data format, Can't have multiple end rooms")
				} else {
					chunks, ok := CheckIsRoom(line, text)
					if ok {
						err := isValidRoom(line, chunks)
						if err != nil {
							return nil, err
						}
						colony.End_room.setRoom(string(chunks[0]), toInt(chunks[1]), toInt(chunks[2]))
						err = colony.addRoom(colony.End_room)
						if err != nil {
							return nil, err
						}
						end_room_found = true
					} else {
						return nil, errors.New("ERROR: invalid data format, Invalid room format at line: " + strconv.Itoa(line))
					}
					end_command_found = false
				}
			} else if !toIgnore(text) && line == colony.End && start_command_found && colony.Start_room.Name == "" {
				return nil, errors.New("ERROR: invalid data format, start room not found")
			} else if !toIgnore(text) && line == colony.Start && end_command_found && colony.End_room.Name == "" {
				return nil, errors.New("ERROR: invalid data format, end room not found")
			} else if chunks, ok := CheckIsRoom(line, text); ok {
				if tunnel_found {
					return nil, errors.New("ERROR: invalid link format, Can't have a room under tunnel section")
				}
				err := isValidRoom(line, chunks)
				if err != nil {
					return nil, err
				}
				new_room := NewRoom()
				new_room.setRoom(string(chunks[0]), toInt(chunks[1]), toInt(chunks[2]))
				err = colony.addRoom(new_room)
				if err != nil {
					return nil, err
				}
			} else if tunnel.Match(text) {
				tunnel_found = true
				err := colony.CheckRooms()
				if err != nil {
					return nil, err
				}
				err = HandleTunnels(colony, text, line)
				if err != nil {
					return nil, err
				}
			} else if _, ok := CheckIsRoom(line, text); !ok {
				if tunnel_found {
					return nil, errors.New("ERROR: invalid format data at line: " + strconv.Itoa(line))
				}
				return nil, errors.New("ERROR: invalid data format, Invalid room format at line: " + strconv.Itoa(line))
			}
		}
	}
	err_struct := colony.CheckStruct()
	if err_struct != nil {
		return nil, err_struct
	}
	return colony, nil
}
