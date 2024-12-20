package main

import (
	"fmt"
	"os"

	"lemin/helpers"
)

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		fmt.Println("USAGE: go run . file.txt")
	case 1:
		colony, err_exe := helpers.Parse(args[0])
		if err_exe != nil {
			helpers.Error(err_exe)
			return
		}
		err_struct := colony.CheckStruct(&colony)
		if err_struct != nil {
			helpers.Error(err_struct)
			return
			
		}
		fmt.Println(colony)
		colony.PrintLinks(colony.Tunnels)

	default:
		fmt.Println("USAGE: go run . file.txt")
	}
}
