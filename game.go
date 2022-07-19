package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Struct that represents the data for one game at one point in time
type Game struct {
	Start_Time time.Time
	End_Time   time.Time

	Home_Team *Team
	Away_Team *Team

	Clock *Clock

	Home_Goals uint8
	Away_Goals uint8
}

func NewGame(home_team, away_team *Team) *Game {
	g := new(Game)
	g.Home_Team = home_team
	g.Away_Team = away_team
	g.Home_Goals = 0
	g.Away_Goals = 0
	g.Clock = NewClock()
	return g
}

func (g *Game) Simulate() {
	home_fwd_line := 0
	away_fwd_line := 0

	fmt.Println("Simulating Game!")

	for g.Clock.Tick() {
		if g.Clock.Seconds_Remaining == 0 {
			home_fwd_line = 0
			away_fwd_line = 0
			continue
		}
		if g.Clock.Seconds_Remaining <= 750 {
			home_fwd_line = 1
			away_fwd_line = 1
		} else if g.Clock.Seconds_Remaining <= 400 {
			home_fwd_line = 2
			away_fwd_line = 2
		} else if g.Clock.Seconds_Remaining <= 150 {
			home_fwd_line = 3
			away_fwd_line = 3
		}

		g.Home_Goals += g.simulateScoring(g.Home_Team, home_fwd_line)
		g.Away_Goals += g.simulateScoring(g.Away_Team, away_fwd_line)
	}

	fmt.Printf("Final Score: %s %d-%d %s\n", g.Home_Team.Abbreviated_Name, g.Home_Goals, g.Away_Goals, g.Away_Team.Abbreviated_Name)
	if g.Home_Goals > g.Away_Goals {
		fmt.Printf("Winner: %s %s\n", g.Home_Team.City, g.Home_Team.Name)
	} else if g.Away_Goals > g.Home_Goals {
		fmt.Printf("Winner: %s %s\n", g.Away_Team.City, g.Away_Team.Name)
	} else {
		fmt.Printf("Tie :(\n")
	}
}

func (g *Game) simulateScoring(t *Team, line int) (goals uint8) {
	goals = 0
	for j := 0; j < 3; j++ {
		player := t.Forwards[line][j]
		p := rand.Float64()
		if p <= player.XGF/(3*PERIOD_LENGTH) {
			fmt.Printf("%s GOAL! Player: %s, %s Time: %d:%d\n",
				t.Abbreviated_Name,
				player.Name.Last,
				player.Name.First,
				g.Clock.Seconds_Remaining/60,
				g.Clock.Seconds_Remaining%60,
			)
			goals++
		}
	}
	return
}
