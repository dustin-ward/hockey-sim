package sim

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
	rand.Seed(time.Now().UnixNano())
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

	fmt.Println("Scoring Summary")
	fmt.Printf("%9s  |  %-25s |  %5s  |  %s\n", "", "PLAYER", "TIME", "PERIOD")
	fmt.Printf("===============================================================\n")

	line_timer := 0
	for g.Clock.Tick() {
		if line_timer == 60 {
			line_timer = 0
			home_fwd_line = (home_fwd_line + 1) % 4
			away_fwd_line = (away_fwd_line + 1) % 4
		}
		line_timer++
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
		if p <= (player.EV_GF60/60)/60 {
			fmt.Printf("%s GOAL!  |  %-25s |  %s  |  %s\n",
				t.Abbreviated_Name,
				fmt.Sprintf("%s, %s", player.Name.Last, player.Name.First),
				fmt.Sprintf("%02d:%02d", g.Clock.Seconds_Remaining/60, g.Clock.Seconds_Remaining%60),
				g.Clock.PrintState(),
			)
			goals++
		}
	}
	return
}
