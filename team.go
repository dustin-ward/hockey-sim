package main

type Position uint8

const (
	LW Position = iota
	C
	RW
	LD
	RD
	G
)

type Team struct {
	Id               uint32
	Name             string
	City             string
	Abbreviated_Name string

	Players []Player

	Forwards [][]*Player
	Defence  [][]*Player
	Goalies  []*Player
}

func NewTeam(name, city, abbreviated_name string) *Team {
	t := new(Team)
	t.Name = name
	t.City = city
	t.Abbreviated_Name = abbreviated_name

	t.Forwards = make([][]*Player, 4)
	for i := 0; i < 4; i++ {
		t.Forwards[i] = make([]*Player, 3)
	}

	t.Defence = make([][]*Player, 3)
	for i := 0; i < 3; i++ {
		t.Defence[i] = make([]*Player, 2)
	}

	t.Goalies = make([]*Player, 2)

	return t
}

func (t *Team) AddPlayer(p Player) {
	t.Players = append(t.Players, p)
}

func (t *Team) RemovePlayer(p Player) {
	for i, v := range t.Players {
		if v.Name == p.Name {
			t.Players = append(t.Players[:i], t.Players[i+1:]...)
			return
		}
	}
}

func (t *Team) ResetLines() {
	for i := 0; i < 4; i++ {
		t.Forwards[i] = make([]*Player, 3)
	}

	for i := 0; i < 3; i++ {
		t.Defence[i] = make([]*Player, 2)
	}

	for i := 0; i < 2; i++ {
		t.Goalies[i] = nil
	}
}

func (t *Team) AssignRole(player *Player, position Position, line int) error {
	switch position {
	case LW, C, RW:
		if line > 2 || line < 0 {
			return INSERT_PLAYER_INVALID_LINE
		}
		t.Forwards[line][position-LW] = player
	case LD, RD:
		if line > 1 || line < 0 {
			return INSERT_PLAYER_INVALID_LINE
		}
		t.Defence[line][position-LD] = player
	case G:
		if line > 1 || line < 0 {
			return INSERT_PLAYER_INVALID_LINE
		}
		t.Goalies[line] = player
	default:
		return INSERT_PLAYER_INVALID_POSITION
	}
	return nil
}

func (t *Team) ValidateLines() bool {
	M := make(map[*Player]bool)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if t.Forwards[i][j] == nil || M[t.Forwards[i][j]] == true {
				return false
			}
			M[t.Forwards[i][j]] = true
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			if t.Defence[i][j] == nil || M[t.Defence[i][j]] == true {
				return false
			}
			M[t.Defence[i][j]] = true
		}
	}

	for i := 0; i < 2; i++ {
		if t.Goalies[i] == nil || M[t.Goalies[i]] == true {
			return false
		}
		M[t.Goalies[i]] = true
	}

	return true
}
