package main

import (
	"fmt"
	"os"

	"lemin/parser"
)

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		fmt.Println("USAGE: go run . file.txt")
	case 1:
		colony, err_exe := parser.Parse(args[0])
		if err_exe != nil {
			parser.Error(err_exe)
			return
		}
		err_struct := colony.CheckStruct(&colony)
		if err_struct != nil {
			parser.Error(err_struct)
			return
			
		}
		fmt.Println(colony)
		colony.PrintLinks(colony.Tunnels)

	default:
		fmt.Println("USAGE: go run . file.txt")
	}
}
