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
		colony, err := helpers.Parse(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(colony.String())
		colony.PrintLinks(colony.Tunnels)

	default:
		fmt.Println("USAGE: go run . file.txt")
	}
}
