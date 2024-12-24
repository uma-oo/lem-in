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
		fmt.Println(colony)
		colony.PrintLinks(colony.Tunnels)
		// p.BFS(colony)
		p.BfsShortestPath(colony , "h")

	default:
		fmt.Println("USAGE: go run . file.txt")
	}
}
