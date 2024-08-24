package tests

import (
	"testing"

	"github.com/Kazyel/Poke-CLI/api"
)

func TestPokedex(t *testing.T) {
	pokedex := api.CreatePokedex()

	pokedex.AddPokemon("Pikachu", "Electric")
	pokedex.AddPokemon("Charmander", "Fire")
	pokedex.AddPokemon("Squirtle", "Water")
	pokedex.AddPokemon("Bulbasaur", "Grass")

	pokemon, _ := pokedex.GetPokemon("Pikachu")

	if pokemon.Name != "Pikachu" {
		t.Errorf("Expected Pokemon name to be 'Pikachu', got '%s'.", pokemon.Name)
	}

	if pokemon.Type != "Electric" {
		t.Errorf("Expected Pokemon type to be 'Electric', got '%s'.", pokemon.Type)
	}
}

func TestPokedexDuplicate(t *testing.T) {
	pokedex := api.CreatePokedex()

	pokedex.AddPokemon("Pikachu", "Electric")
	pokedex.AddPokemon("Charmander", "Fire")
	pokedex.AddPokemon("Squirtle", "Water")
	pokedex.AddPokemon("Bulbasaur", "Grass")

	err := pokedex.AddPokemon("Pikachu", "Electric")

	if err == nil {
		t.Errorf("Expected to not be able to add duplicate Pokemon.")
		return
	}
}

func TestPokedexDeletion(t *testing.T) {
	pokedex := api.CreatePokedex()

	pokedex.AddPokemon("Pikachu", "Electric")
	pokedex.AddPokemon("Charmander", "Fire")
	pokedex.AddPokemon("Squirtle", "Water")
	pokedex.AddPokemon("Bulbasaur", "Grass")

	err := pokedex.WithdrawPokemon("Pikachu")

	if err != nil {
		t.Errorf("Expected to be able to withdraw Pokemon.")
		return
	}
}
