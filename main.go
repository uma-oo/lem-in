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
		fmt.Println("USAGE: go run . file.txt")
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
		fmt.Println(colony.Bad_Rooms)                                                              // Detected here before the Start of the BFS
		colony.Shortest_Path = p.BaseBFS(colony, colony.Start_room.Name, colony.End_room.Name)[1:] // Put the ShortestPath into the colony
		fmt.Println("length: ", len(colony.Shortest_Path), "The Shortest Path Found: ", colony.Shortest_Path)
		// p.BFS(colony)
		// p.BfsShortestPath(colony , "h", "end")

		// fmt.Println(p.DegreeNeighbors(colony))
		// fmt.Println(p.Priority(colony))
		// fmt.Println(p.RunnerBFS(colony))
		paths := p.BFSOptimizedDisjoint(colony, colony.Start_room.Name, colony.End_room.Name)
		fmt.Println("=====================>", len(paths))
		for _, path := range paths {
			fmt.Println("Length: ", len(path), "path: ", path)
		}

		fmt.Println(p.AverageRoomLinks(colony))
		p.DecideWhichPath(colony)
		fmt.Println(p.PriorityWithBadRoom(colony))
		combianaison:=p.BFSCombinaisons(colony, colony.Start_room.Name, colony.End_room.Name)
		for _,ele:= range combianaison {
			fmt.Println(ele)
		}

	default:
		fmt.Println("USAGE: go run . file.txt")
	}
}
