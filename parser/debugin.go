package parser

import "fmt"

func PrintGroupsFound(groups []*Group) {
	for index, grp := range groups {
		fmt.Printf("grp[%v]:\n", index)
		for _, path := range grp.Paths {
			fmt.Printf("path.Rooms_found: %v\n", path.Rooms_found)
		}
	}
}
