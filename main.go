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
		err_struct := colony.CheckStruct(&colony)
		if err_struct != nil {
			p.Error(err_struct)
			return

		}

		groups2 := p.FindAllGroups(colony)
		for _, g := range groups2 {
			fmt.Printf("group: %v\n", g.String())
		}
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
