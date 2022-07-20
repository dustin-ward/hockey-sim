package sim

func Oilers() *Team {
	Oilers := NewTeam("Oilers", "Edmonton", "EDM")
	Oilers.AssignRole(&Player{Name: Name{Last: "McDavid", First: "Connor"}, XGF: 1.44}, C, 0)
	Oilers.AssignRole(&Player{Name: Name{Last: "Kane", First: "Evander"}, XGF: 1.33}, LW, 0)
	Oilers.AssignRole(&Player{Name: Name{Last: "Hyman", First: "Zach"}, XGF: 1.39}, RW, 0)
	Oilers.AssignRole(&Player{Name: Name{Last: "Draisaitl", First: "Leon"}, XGF: 1.43}, C, 1)
	Oilers.AssignRole(&Player{Name: Name{Last: "Nugent-Hopkins", First: "Ryan"}, XGF: 0.75}, LW, 1)
	Oilers.AssignRole(&Player{Name: Name{Last: "Yamamoto", First: "Kailer"}, XGF: 0.82}, RW, 1)
	Oilers.AssignRole(&Player{Name: Name{Last: "McLeod", First: "Ryan"}, XGF: 0.56}, C, 2)
	Oilers.AssignRole(&Player{Name: Name{Last: "Foegele", First: "Warren"}, XGF: 1.01}, LW, 2)
	Oilers.AssignRole(&Player{Name: Name{Last: "Puljujarvi", First: "Jesse"}, XGF: 1.19}, RW, 2)
	Oilers.AssignRole(&Player{Name: Name{Last: "Ryan", First: "Derek"}, XGF: 0.62}, C, 3)
	Oilers.AssignRole(&Player{Name: Name{Last: "Shore", First: "Devin"}, XGF: 0.52}, LW, 3)
	Oilers.AssignRole(&Player{Name: Name{Last: "Kassian", First: "Zach"}, XGF: 0.64}, RW, 3)

	Oilers.AssignRole(&Player{Name: Name{Last: "Nurse", First: "Darnell"}, XGF: 0.36}, LD, 0)
	Oilers.AssignRole(&Player{Name: Name{Last: "Ceci", First: "Cody"}, XGF: 0.12}, RD, 0)
	Oilers.AssignRole(&Player{Name: Name{Last: "Keith", First: "Duncan"}, XGF: 0.15}, LD, 1)
	Oilers.AssignRole(&Player{Name: Name{Last: "Bouchard", First: "Evan"}, XGF: 0.3}, RD, 1)
	Oilers.AssignRole(&Player{Name: Name{Last: "Kulak", First: "Brett"}, XGF: 0.13}, LD, 2)
	Oilers.AssignRole(&Player{Name: Name{Last: "Barrie", First: "Tyson"}, XGF: 0.26}, RD, 2)

	return Oilers
}

func Sabres() *Team {
	Sabres := NewTeam("Sabres", "Buffalo", "BUF")
	Sabres.AssignRole(&Player{Name: Name{Last: "Thompson", First: "Tage"}, XGF: 1.12}, C, 0)
	Sabres.AssignRole(&Player{Name: Name{Last: "Skinner", First: "Jeff"}, XGF: 1.45}, LW, 0)
	Sabres.AssignRole(&Player{Name: Name{Last: "Olofsson", First: "Victor"}, XGF: 0.9}, RW, 0)
	Sabres.AssignRole(&Player{Name: Name{Last: "Cozens", First: "Dylan"}, XGF: 0.77}, C, 1)
	Sabres.AssignRole(&Player{Name: Name{Last: "Tuch", First: "Alex"}, XGF: 1.07}, LW, 1)
	Sabres.AssignRole(&Player{Name: Name{Last: "Mittelstadt", First: "Casey"}, XGF: 0.65}, RW, 1)
	Sabres.AssignRole(&Player{Name: Name{Last: "Girgensons", First: "Zemgus"}, XGF: 0.64}, C, 2)
	Sabres.AssignRole(&Player{Name: Name{Last: "Hinostroza", First: "Vinnie"}, XGF: 0.58}, LW, 2)
	Sabres.AssignRole(&Player{Name: Name{Last: "Asplund", First: "Rasmus"}, XGF: 0.76}, RW, 2)
	Sabres.AssignRole(&Player{Name: Name{Last: "Hayden", First: "John"}, XGF: 0.54}, C, 3)
	Sabres.AssignRole(&Player{Name: Name{Last: "Bjork", First: "Anders"}, XGF: 0.35}, LW, 3)
	Sabres.AssignRole(&Player{Name: Name{Last: "Krebs", First: "Peyton"}, XGF: 0.35}, RW, 3)

	Sabres.AssignRole(&Player{Name: Name{Last: "Girgensons", First: "Zemgus"}, XGF: 0.64}, LD, 0)
	Sabres.AssignRole(&Player{Name: Name{Last: "Hinostroza", First: "Vinnie"}, XGF: 0.58}, RD, 0)
	Sabres.AssignRole(&Player{Name: Name{Last: "Asplund", First: "Rasmus"}, XGF: 0.76}, LD, 1)
	Sabres.AssignRole(&Player{Name: Name{Last: "Hayden", First: "John"}, XGF: 0.54}, RD, 1)
	Sabres.AssignRole(&Player{Name: Name{Last: "Bjork", First: "Anders"}, XGF: 0.35}, LD, 2)
	Sabres.AssignRole(&Player{Name: Name{Last: "Krebs", First: "Peyton"}, XGF: 0.35}, RD, 2)

	return Sabres
}

func Lightning() *Team {
	Lightning := NewTeam("Lightning", "Tampa Bay", "TBL")
	Lightning.AssignRole(&Player{Name: Name{Last: "Stamkos", First: "Steven"}, XGF: 1.07}, C, 0)
	Lightning.AssignRole(&Player{Name: Name{Last: "Palat", First: "Ondrej"}, XGF: 0.87}, LW, 0)
	Lightning.AssignRole(&Player{Name: Name{Last: "Kucherov", First: "Nikita"}, XGF: 1.05}, RW, 0)
	Lightning.AssignRole(&Player{Name: Name{Last: "Point", First: "Brayden"}, XGF: 1.23}, C, 1)
	Lightning.AssignRole(&Player{Name: Name{Last: "Paul", First: "Nick"}, XGF: 0.85}, LW, 1)
	Lightning.AssignRole(&Player{Name: Name{Last: "Colton", First: "Ross"}, XGF: 1.15}, RW, 1)
	Lightning.AssignRole(&Player{Name: Name{Last: "Cirelli", First: "Anthony"}, XGF: 0.79}, C, 2)
	Lightning.AssignRole(&Player{Name: Name{Last: "Killorn", First: "Alex"}, XGF: 0.91}, LW, 2)
	Lightning.AssignRole(&Player{Name: Name{Last: "Hagel", First: "Brandon"}, XGF: 0.85}, RW, 2)
	Lightning.AssignRole(&Player{Name: Name{Last: "Bellemare", First: "Pierre-Edouard"}, XGF: 0.49}, C, 3)
	Lightning.AssignRole(&Player{Name: Name{Last: "Maroon", First: "Pat"}, XGF: 0.77}, LW, 3)
	Lightning.AssignRole(&Player{Name: Name{Last: "Perry", First: "Corey"}, XGF: 1.36}, RW, 3)

	Lightning.AssignRole(&Player{Name: Name{Last: "Cirelli", First: "Anthony"}, XGF: 0.79}, LD, 0)
	Lightning.AssignRole(&Player{Name: Name{Last: "Killorn", First: "Alex"}, XGF: 0.91}, RD, 0)
	Lightning.AssignRole(&Player{Name: Name{Last: "Hagel", First: "Brandon"}, XGF: 0.85}, LD, 1)
	Lightning.AssignRole(&Player{Name: Name{Last: "Bellemare", First: "Pierre-Edouard"}, XGF: 0.49}, RD, 1)
	Lightning.AssignRole(&Player{Name: Name{Last: "Maroon", First: "Pat"}, XGF: 0.77}, LD, 2)
	Lightning.AssignRole(&Player{Name: Name{Last: "Perry", First: "Corey"}, XGF: 1.36}, RD, 2)

	return Lightning
}
