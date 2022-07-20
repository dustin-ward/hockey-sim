package main

import "dustin-ward/GMSim/sim"

func main() {
	G := sim.NewGame(sim.Sabres(), sim.Lightning())
	G.Simulate()
}
