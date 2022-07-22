package sim

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Position uint8

const (
	LW Position = iota
	C
	RW
	LD
	RD
	G

	F
	D
	CLW
	CRW

	INVALID
)

func GetPositionEnum(pos string) (Position, error) {
	switch pos {
	case "LW":
		return LW, nil
	case "C":
		return C, nil
	case "RW":
		return RW, nil
	case "LD":
		return LD, nil
	case "RD":
		return RD, nil
	case "G":
		return G, nil
	case "F":
		return F, nil
	case "D":
		return D, nil
	case "CLW":
		return CLW, nil
	case "CRW":
		return CRW, nil
	default:
		return INVALID, INVALID_POSITION
	}
}

type Team struct {
	Id               uint32
	Name             string
	City             string
	Abbreviated_Name string

	Players []*Player

	Forwards [][]*Player
	Defence  [][]*Player
	Goalies  []*Player
}

func NewTeam(name, city, abbreviated_name, year string) (*Team, error) {
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

	if err := t.ReadTeamFromCsv(year); err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Team) AddPlayer(p *Player) {
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

func (t *Team) GetPlayer(name Name) *Player {
	for _, p := range t.Players {
		if p.Name == name {
			return p
		}
	}
	return nil
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
		if line > 3 || line < 0 {
			return INSERT_PLAYER_INVALID_LINE
		}
		t.Forwards[line][position-LW] = player
	case LD, RD:
		if line > 2 || line < 0 {
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

func (t *Team) ReadTeamFromCsv(year string) error {
	lowered_name := strings.ToLower(t.Name)
	f, err := os.Open(fmt.Sprintf("./data/%s_players_%s.csv", lowered_name, year))
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		return err
	}

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {
		p, err := NewPlayerFromSlice(record)
		if err != nil {
			return err
		}
		t.Players = append(t.Players, p)
	}

	return nil
}
