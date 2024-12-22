package parser

func NewColony() Colony {
	return Colony{
		Ants:       0,
		Start:      0,
		End:        0,
		Start_room: NewRoom(),
		End_room:   NewRoom(),
		Rooms_coor: make(map[string][]interface{}),
		Tunnels:    make(map[string]*Room),
	}
}

func NewRoom() *Room {
	return &Room{
		Links: map[string]struct{}{},
	}
}
