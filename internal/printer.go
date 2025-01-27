package internal

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func (G *Group) InitializeMvt(graph *Colony) []*Agent {
	var Agents map[int]*Agent = make(map[int]*Agent)
	var Agents_slice []*Agent
	// This is the correct version of the function and nothing more or less
	for i := 1; i <= graph.Ants; i++ {
		if _, ok := Agents[i]; !ok {
			new_ant := NewAgent()
			shortest := new_ant.FindPath(G, Agents)
			new_ant.CountPath(Agents_slice)
			shortest.Length += 1
			Agents[i] = new_ant
			Agents_slice = append(Agents_slice, new_ant)
		}
	}
	return Agents_slice
}

func (g *Group) MoveAnts(graph *Colony) []string {
	lines := []string{}
	agents := g.InitializeMvt(graph) // Find the paths for each ant in the colony
	for i := 1; i <= g.Turns; i++ {
		line := ""
		positions := make(map[string]struct{})
		for j, ant := range agents {
			pos := GetAntPos(i, ant.Pos, ant.PathUsed.Rooms_found)
			_, ok := positions[pos]
			// fmt.Printf("ant: %v pos: %v turn: %v\n", ant, pos, i)
			if pos != "" && !ok {
				line += fmt.Sprintf("L%v-%v ", j+1, pos)
				if pos != graph.End_room.Name {
					positions[pos] = struct{}{}
				}
			}
		}
		fmt.Println(line)
	}
	return lines
}

// given a specific ant and turn we get l position dyal ant f l path
func GetAntPos(turn int, ant int, path []string) string {
	if turn-ant < len(path) && turn-ant >= 0 {
		return path[turn-ant]
	}
	return ""
}

func (A *Agent) FindPath(group_chosen *Group, agents map[int]*Agent) *Path {
	// Find The Shortest Path inside The group and then assign it to the ant
	shortest_path := group_chosen.ReturnShortestPath()
	A.PathUsed = shortest_path
	return shortest_path
}

// The idea is as follows
// if the path has been taken by another ant meaning by this, it's not her first time to appear
// we index the ant using the Pos and the Pos means only that it's the first the second or the third in the path
// like having a primary key (path , pos) pos reflects the turn for that specific path

func (A *Agent) CountPath(agents []*Agent) {
	count := 1
	for _, agent := range agents {
		if reflect.DeepEqual(A.PathUsed.Rooms_found, agent.PathUsed.Rooms_found) {
			count++
		}
	}
	A.Pos = count
}

func ReadFile(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	defer file.Close()
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
