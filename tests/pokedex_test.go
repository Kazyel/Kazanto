package tests

import (
	"testing"

	"github.com/Kazyel/Poke-CLI/api"
)

func TestPokedex(t *testing.T) {
	pokedex := api.CreatePokedex()

	api.AddPokemon("Pikachu", "Electric", pokedex)
	api.AddPokemon("Charmander", "Fire", pokedex)
	api.AddPokemon("Squirtle", "Water", pokedex)
	api.AddPokemon("Bulbasaur", "Grass", pokedex)

	pokemon, _ := api.GetPokemon("Pikachu", pokedex)

	if pokemon.Name != "Pikachu" {
		t.Errorf("Expected Pokemon name to be 'Pikachu', got '%s'.", pokemon.Name)
	}

	if pokemon.Type != "Electric" {
		t.Errorf("Expected Pokemon type to be 'Electric', got '%s'.", pokemon.Type)
	}
}

func TestPokedexDuplicate(t *testing.T) {
	pokedex := api.CreatePokedex()

	api.AddPokemon("Pikachu", "Electric", pokedex)
	api.AddPokemon("Charmander", "Fire", pokedex)
	api.AddPokemon("Squirtle", "Water", pokedex)
	api.AddPokemon("Bulbasaur", "Grass", pokedex)

	err := api.AddPokemon("Pikachu", "Electric", pokedex)

	if err == nil {
		t.Errorf("Expected to not be able to add duplicate Pokemon.")
		return
	}
}

func TestPokedexDeletion(t *testing.T) {
	pokedex := api.CreatePokedex()

	api.AddPokemon("Pikachu", "Electric", pokedex)
	api.AddPokemon("Charmander", "Fire", pokedex)
	api.AddPokemon("Squirtle", "Water", pokedex)
	api.AddPokemon("Bulbasaur", "Grass", pokedex)

	err := api.WithdrawPokemon("Pikachu", pokedex)

	if err != nil {
		t.Errorf("Expected to be able to withdraw Pokemon.")
		return
	}
}
