package main

import (
	"fmt"
	"os"
	"strings"

	p "lemin/parser"
)

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 1:
		colony, err_exe := p.Parse(args[0])
		if err_exe != nil {
			p.Error(err_exe)
			return
		}
		err_struct := colony.CheckStruct()
		if err_struct != nil {
			p.Error(err_struct)
			return
		}
		p.ReadFile(args[0])
		fmt.Print("\n\n")
		// colony.PrintLinks(colony.Tunnels)
		// fmt.Printf("Levels %v\n", p.Levels(colony, colony.Start_room.Name, colony.End_room.Name))
		// groups := p.FindAllGroups(colony)
		// p.PrintGroupsFound(groups)
		// for _,group := range groups {
		// 	fmt.Println(group.String())
		// 	group.CalculTurns(colony)
		// 	fmt.Println(group.Turns)
		// }
		best := p.FindTheBestGrp(colony)
		// fmt.Println(best.String())
		// best.CalculTurns(colony)
		best.MoveAnts(colony)

	default:
		fmt.Println("USAGE: go run . file.txt")
	}
}

func init() {
	args := os.Args[1:]
	switch true {
	case len(args) != 1:
		p.Error("USAGE: go run . file.txt")
		os.Exit(0)
	case !strings.HasSuffix(args[0], ".txt"):
		p.Error("USAGE: go run . file.txt")
		os.Exit(0)
	}
}
