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
		fmt.Println(colony.Bad_Rooms)// Detected here before the Start of the BFS 
		shortest_path := p.BaseBFS(colony, "t", colony.End_room.Name)
		fmt.Println("length: ",len(shortest_path)-1,"The Shortest Path Found: ", shortest_path)
		// p.BFS(colony)
		// p.BfsShortestPath(colony , "h", "end")
		// fmt.Println(len(colony.Tunnels))
		// fmt.Println(p.DegreeNeighbors(colony))
		// fmt.Println(p.Priority(colony))
		// fmt.Println(p.RunnerBFS(colony))
		paths := p.BFSOptimized(colony, colony.Start_room.Name, colony.End_room.Name)
		fmt.Println("=====================>", len(paths))
		for _, path := range paths {
			fmt.Println("Length: ",len(path),"path: ", path)
		}


		fmt.Println(p.AverageRoomLinks(colony))

	default:
		fmt.Println("USAGE: go run . file.txt")
	}
}
