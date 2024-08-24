package api

import "fmt"

type Pokemon struct {
	Name string
	Type string
}

type Pokedex struct {
	Pokemons map[string]Pokemon
}

func CreatePokedex() *Pokedex {
	return &Pokedex{
		Pokemons: map[string]Pokemon{},
	}
}

func (pokedex *Pokedex) GetPokedex() map[string]Pokemon {
	return pokedex.Pokemons
}

func (pokedex *Pokedex) GetPokemon(name string) (Pokemon, error) {
	if pokedex.Pokemons[name].Name == "" {
		fmt.Println("Pokemon not captured.")
		return Pokemon{}, fmt.Errorf("Pokemon not captured")
	}

	return pokedex.Pokemons[name], nil
}

func (pokedex *Pokedex) AddPokemon(name string, typeName string) error {
	if pokedex.Pokemons[name].Name != "" {
		fmt.Println("Pokemon already captured.")
		return fmt.Errorf("Pokemon already captured")
	}

	pokedex.Pokemons[name] = Pokemon{
		Name: name,
		Type: typeName,
	}

	return nil
}

func (pokedex *Pokedex) WithdrawPokemon(name string) error {
	if pokedex.Pokemons[name].Name == "" {
		fmt.Println("Pokemon not captured.")
		return fmt.Errorf("Pokemon not captured")
	}

	delete(pokedex.Pokemons, name)
	return nil
}
