// this is just for debuging purposes
package internal

import "fmt"

func (c *Colony) String() string {
	return fmt.Sprintf("Colony(Number of ants: %v, Start: %v, End: %v, Start Room: %v, End Room: %v , Tunnels:%v)", c.Ants, c.Start, c.End, c.Start_room.String(), c.End_room.String(), c.Tunnels)
}

func (r *Room) String() string {
	return fmt.Sprintf("Room %v and its links are %v", r.Name, r.Links)
}

func (a *Agent) String() string {
	return fmt.Sprintf("Ant Position in the path: %v, Path used: %v", a.Pos, a.PathUsed.String())
}

func (t *Traversal) String() string {
	return fmt.Sprintf("Visited_nodes: %v queue: %v", t.Visited_Node, t.Queue)
}

func (c *Colony) PrintLinks(links map[string]*Room) {
	for key, value := range links {
		fmt.Printf("the room  %s and the links are %v \n", key, value.Links)
	}
}

func (P *Path) String() []string {
	return P.Rooms_found
}

func (G *Group) String() [][]string {
	paths := [][]string{}
	for _, path := range G.Paths {
		paths = append(paths, path.Rooms_found)
	}
	return paths
}
