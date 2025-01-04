package main

import (
	"fmt"
	"os"

	p "lemin/parser"
)

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		p.Error("USAGE: go run . file.txt")
		return
	case 1:
		colony, err_exe := p.Parse(args[0])
		if err_exe != nil {
			p.Error(err_exe)
			return
		}
		err_struct := colony.CheckStruct(&colony)
		if err_struct != nil {
			p.Error(err_struct)
			return

		}
		p.DetectBadRooms(colony)

		colony.Shortest_Path = p.BFS(colony, colony.Start_room.Name, colony.End_room.Name)[1:] // Put the ShortestPath into the colony
		fmt.Println("length: ", len(colony.Shortest_Path), "The Shortest Path Found: ", colony.Shortest_Path)
		groups := p.RunnerBFS(colony)
		for _, group := range groups {
			fmt.Println(group.String())
		}

	default:
		fmt.Println("USAGE: go run . file.txt")
	}
}

// func init() {
// 	args := os.Args[1:]
// 	switch true {
// 	case len(args) != 1:
// 		p.Error("USAGE: go run . file.txt")
// 		return
// 	case !strings.HasSuffix(args[0], ".txt"):
// 		p.Error("USAGE: go run . file.txt")
// 		return

// 	}
// }
