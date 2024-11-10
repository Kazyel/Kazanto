package api

type Pokeball struct {
	Name            string
	Description     string
	CatchMultiplier float64
}

type Item struct {
	Name     string
	Quantity int
}

func createPokeball() Pokeball {
	return Pokeball{
		Name:            "Pokeball",
		Description:     "A device that allows you to catch wild Pokémons.",
		CatchMultiplier: 1.0,
	}
}

func createMasterBall() Pokeball {
	return Pokeball{
		Name:            "Master Ball",
		Description:     "A device that allows you to catch wild Pokémons.",
		CatchMultiplier: 100.0,
	}
}

func createUltraBall() Pokeball {
	return Pokeball{
		Name:            "Ultra Ball",
		Description:     "A device that allows you to catch wild Pokémons.",
		CatchMultiplier: 2.0,
	}
}
