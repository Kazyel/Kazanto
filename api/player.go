package api

import (
	"github.com/Kazyel/Poke-CLI/utils"
)

type Player struct {
	Name      string
	Pokedex   *Pokedex
	Inventory *Inventory
	Party     []Pokemon
}

func CreateInventory() *Inventory {
	return &Inventory{
		Pokeballs: map[string]*Pokeballs{
			"pokeball": {
				Pokeball: createPokeball(),
				Quantity: 5,
			},
			"masterball": {
				Pokeball: createMasterBall(),
				Quantity: 5,
			},
			"ultraball": {
				Pokeball: createUltraBall(),
				Quantity: 5,
			},
		},
	}
}

func CreateNewPlayer(name string) *Player {
	return &Player{
		Name:      name,
		Pokedex:   CreatePokedex(),
		Inventory: CreateInventory(),
		Party:     []Pokemon{},
	}
}

func (player *Player) InspectInventory() {
	utils.PrintTitle("Your Pokéballs:")
	player.RenderInventory()
}
