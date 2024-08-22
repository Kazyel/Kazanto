package api

import "fmt"

type Pokemon struct {
	Name string
	Type string
}

type Pokedex struct {
	Pokemons map[string]Pokemon
}

func GetPokedex() string {
	return "pokedex"
}

func AddPokemon(name string, typeName string, pokedex *Pokedex) error {
	if pokedex.Pokemons[name].Name != "" {
		fmt.Println("Pokemon already captured.")
		return nil
	}

	pokedex.Pokemons[name] = Pokemon{
		Name: name,
		Type: typeName,
	}

	return nil
}

func GetPokemon(name string, pokedex *Pokedex) Pokemon {
	return pokedex.Pokemons[name]
}

func WithdrawPokemon(name string, pokedex *Pokedex) {
	delete(pokedex.Pokemons, name)
}
