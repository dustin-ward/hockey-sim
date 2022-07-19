package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	Team1 := NewTeam("Oilers", "Edmonton", "EDM")
	Team2 := NewTeam("Sabres", "Buffalo", "BUF")

	Team1.AssignRole(&Player{Name: Name{Last: "McDavid", First: "Connor"}, XGF: 1.44}, C, 0)
	Team1.AssignRole(&Player{Name: Name{Last: "Kane", First: "Evander"}, XGF: 1.33}, LW, 0)
	Team1.AssignRole(&Player{Name: Name{Last: "Hyman", First: "Zach"}, XGF: 1.39}, RW, 0)
	Team1.AssignRole(&Player{Name: Name{Last: "Draisaitl", First: "Leon"}, XGF: 1.43}, C, 1)
	Team1.AssignRole(&Player{Name: Name{Last: "Nugent-Hopkins", First: "Ryan"}, XGF: 0.75}, LW, 1)
	Team1.AssignRole(&Player{Name: Name{Last: "Yamamoto", First: "Kailer"}, XGF: 0.82}, RW, 1)
	Team1.AssignRole(&Player{Name: Name{Last: "McLeod", First: "Ryan"}, XGF: 0.56}, C, 2)
	Team1.AssignRole(&Player{Name: Name{Last: "Foegele", First: "Warren"}, XGF: 1.01}, LW, 2)
	Team1.AssignRole(&Player{Name: Name{Last: "Puljujarvi", First: "Jesse"}, XGF: 1.19}, RW, 2)
	Team1.AssignRole(&Player{Name: Name{Last: "Ryan", First: "Derek"}, XGF: 0.62}, C, 3)
	Team1.AssignRole(&Player{Name: Name{Last: "Shore", First: "Devin"}, XGF: 0.52}, LW, 3)
	Team1.AssignRole(&Player{Name: Name{Last: "Kassian", First: "Zach"}, XGF: 0.64}, RW, 3)

	Team2.AssignRole(&Player{Name: Name{Last: "Thompson", First: "Tage"}, XGF: 1.12}, C, 0)
	Team2.AssignRole(&Player{Name: Name{Last: "Skinner", First: "Jeff"}, XGF: 1.45}, LW, 0)
	Team2.AssignRole(&Player{Name: Name{Last: "Olofsson", First: "Victor"}, XGF: 0.9}, RW, 0)
	Team2.AssignRole(&Player{Name: Name{Last: "Cozens", First: "Dylan"}, XGF: 0.77}, C, 1)
	Team2.AssignRole(&Player{Name: Name{Last: "Tuch", First: "Alex"}, XGF: 1.07}, LW, 1)
	Team2.AssignRole(&Player{Name: Name{Last: "Mittelstadt", First: "Casey"}, XGF: 0.65}, RW, 1)
	Team2.AssignRole(&Player{Name: Name{Last: "Girgensons", First: "Zemgus"}, XGF: 0.64}, C, 2)
	Team2.AssignRole(&Player{Name: Name{Last: "Hinostroza", First: "Vinnie"}, XGF: 0.58}, LW, 2)
	Team2.AssignRole(&Player{Name: Name{Last: "Asplund", First: "Rasmus"}, XGF: 0.76}, RW, 2)
	Team2.AssignRole(&Player{Name: Name{Last: "Hayden", First: "John"}, XGF: 0.54}, C, 3)
	Team2.AssignRole(&Player{Name: Name{Last: "Bjork", First: "Anders"}, XGF: 0.35}, LW, 3)
	Team2.AssignRole(&Player{Name: Name{Last: "Krebs", First: "Peyton"}, XGF: 0.35}, RW, 3)

	G := NewGame(Team1, Team2)
	G.Simulate()
}
