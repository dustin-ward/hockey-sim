package main

import (
	"dustin-ward/GMSim/sim"
	"log"
)

func main() {
	Oilers, err := sim.NewTeam("Oilers", "Edmonton", "EDM", "2021")
	if err != nil {
		log.Fatal(err)
	}
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "McDavid", First: "Connor"}), sim.C, 0)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Kane", First: "Evander"}), sim.LW, 0)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Hyman", First: "Zach"}), sim.RW, 0)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Draisaitl", First: "Leon"}), sim.C, 1)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Nugent-Hopkins", First: "Ryan"}), sim.LW, 1)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Yamamoto", First: "Kailer"}), sim.RW, 1)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "McLeod", First: "Ryan"}), sim.C, 2)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Foegele", First: "Warren"}), sim.LW, 2)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Puljujarvi", First: "Jesse"}), sim.RW, 2)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Ryan", First: "Derek"}), sim.C, 3)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Shore", First: "Devin"}), sim.LW, 3)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Kassian", First: "Zack"}), sim.RW, 3)

	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Nurse", First: "Darnell"}), sim.LD, 0)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Ceci", First: "Cody"}), sim.RD, 0)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Keith", First: "Duncan"}), sim.LD, 1)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Bouchard", First: "Evan"}), sim.RD, 1)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Kulak", First: "Brett"}), sim.LD, 2)
	Oilers.AssignRole(Oilers.GetPlayer(sim.Name{Last: "Barrie", First: "Tyson"}), sim.RD, 2)

	Sabres, err := sim.NewTeam("Sabres", "Buffalo", "BUF", "2021")
	if err != nil {
		log.Fatal(err)
	}
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Thompson", First: "Tage"}), sim.C, 0)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Skinner", First: "Jeff"}), sim.LW, 0)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Olofsson", First: "Victor"}), sim.RW, 0)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Cozens", First: "Dylan"}), sim.C, 1)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Tuch", First: "Alex"}), sim.LW, 1)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Mittelstadt", First: "Casey"}), sim.RW, 1)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Girgensons", First: "Zemgus"}), sim.C, 2)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Hinostroza", First: "Vinnie"}), sim.LW, 2)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Asplund", First: "Rasmus"}), sim.RW, 2)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Hayden", First: "John"}), sim.C, 3)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Bjork", First: "Anders"}), sim.LW, 3)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Krebs", First: "Peyton"}), sim.RW, 3)

	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Girgensons", First: "Zemgus"}), sim.LD, 0)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Hinostroza", First: "Vinnie"}), sim.RD, 0)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Asplund", First: "Rasmus"}), sim.LD, 1)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Hayden", First: "John"}), sim.RD, 1)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Bjork", First: "Anders"}), sim.LD, 2)
	Sabres.AssignRole(Sabres.GetPlayer(sim.Name{Last: "Krebs", First: "Peyton"}), sim.RD, 2)

	G := sim.NewGame(Oilers, Sabres)
	G.Simulate()
}
