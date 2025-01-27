package main

import (
	"fmt"
	"os"
	"strings"

	p "lemin/internal"
)

func main() {
	args := os.Args[1:]
	colony, err_exe := p.Parse(args[0])
	if err_exe != nil {
		p.Error(err_exe)
		return
	}
	bestGroup := p.FindTheBestGrp(colony)
	if len(bestGroup.Paths) != 0 {
		p.ReadFile(args[0])
		fmt.Println()
		bestGroup.MoveAnts(colony)
	} else {
		p.Error("ERROR: No paths have been found.")
		return
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
