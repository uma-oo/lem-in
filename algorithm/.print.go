package utils

import (
	"fmt"
	"slices"

	"lem-in/utils/parser"
)

type PathTracker struct {
	PathGroups []*PathGroup
}

type PathGroup struct {
	Paths     map[string]Path
	UsedRooms []string
}

type Path struct {
	Rooms  []string
	Length int
}

func (pg *PathGroup) FindShortestPath(start string, af *parser.AntFarm) {
	queue := []string{start}
	previousRoomTracker := map[string]string{}
	// previousRoomTracker[start] = af.StartRoom
	isVisisted := map[string]bool{}
	isVisisted[start] = true

	for len(queue) > 0 {
		currentRoom := queue[0]
		queue = queue[1:]
		for neighbor := range af.Rooms[currentRoom].Links {
			if !slices.Contains(pg.UsedRooms, neighbor) && !isVisisted[neighbor] {
				isVisisted[neighbor] = true
				queue = append(queue, neighbor)
				previousRoomTracker[neighbor] = currentRoom
				if neighbor == af.EndRoom {
					pg.Paths[start] = buildPath(previousRoomTracker, af.EndRoom)
					lenPath := pg.Paths[start].Length
					pg.UsedRooms = append(pg.UsedRooms, pg.Paths[start].Rooms[:lenPath-1]...)
					return
				}
			}
		}
	}
}
func FindPaths(af *parser.AntFarm) *PathTracker {
	roomsLinkedtoStart := extractRoomsLinkedToStart(af.Rooms[af.StartRoom].Links)
	pathTracker := PathTracker{}
	for _, prioritizedRoom := range roomsLinkedtoStart {
		pathGroup := &PathGroup{
			UsedRooms: []string{af.StartRoom},
			Paths:     map[string]Path{},
		}
		// Find path through prioritized room first
		pathGroup.FindShortestPath(prioritizedRoom, af)
		// fmt.Printf("Prioritized room %v => %v\n", prioritizedRoom, pathGroup.Paths[prioritizedRoom].Rooms)
		// fmt.Printf("pathGroup.UsedRooms: %v\n", pathGroup.UsedRooms)
		// fmt.Println(pathGroup.Paths)
		// Then try other rooms that haven't been used yet
		for _, room := range roomsLinkedtoStart {
			if prioritizedRoom != room {
				pathGroup.FindShortestPath(room, af)
				// fmt.Printf("pathGroup.UsedRooms: %v\n", pathGroup.UsedRooms)
			}
		}
		pathTracker.PathGroups = append(pathTracker.PathGroups, pathGroup)
	}
	for i, grp := range pathTracker.PathGroups {
		fmt.Println("grp ", i)
		for startRoom, path := range grp.Paths {
			fmt.Printf("startRoom: %v ==> %v\n", startRoom, path.Rooms)
		}
	}
	return &pathTracker
}

func extractRoomsLinkedToStart(rooms map[string]struct{}) []string {
	result := []string{}
	for room := range rooms {
		result = append(result, room)
	}
	return result
}

func buildPath(previousRoomTracker map[string]string, endRoom string) Path {
	rooms := []string{}
	for currentRoom := endRoom; currentRoom != ""; currentRoom = previousRoomTracker[currentRoom] {
		rooms = append([]string{currentRoom}, rooms...)
	}

	return Path{
		Rooms:  rooms,
		Length: len(rooms),
	}
}