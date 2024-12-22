package parser

func NewColony() colony {
	return colony{
		ants:       0,
		start:      0,
		end:        0,
		start_room: NewRoom(),
		end_room:   NewRoom(),
		rooms_coor: make(map[string][]interface{}),
		Tunnels:    make(map[string]*room),
	}
}

func NewRoom() *room {
	return &room{
		Links: map[string]struct{}{},
	}
}
