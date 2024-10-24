package api

import (
	"fmt"

	"github.com/Kazyel/Poke-CLI/utils"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

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

func (pokedex *Pokedex) RenderPokedex() {
	headerFmt := color.New(color.FgHiBlue, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Name", "Type")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, pokemon := range pokedex.Pokemons {
		tbl.AddRow(pokemon.Name, pokemon.Type)
	}

	tbl.Print()
}

func (pokedex *Pokedex) GetPokemon(name string) (Pokemon, error) {
	if pokedex.Pokemons[name].Name == "" {
		utils.PrintError("Pokemon not captured.")
		return Pokemon{}, fmt.Errorf("Pokemon not captured")
	}

	return pokedex.Pokemons[name], nil
}

func (pokedex *Pokedex) AddPokemon(name string, typeName string) error {
	if pokedex.Pokemons[name].Name != "" {
		utils.PrintError("Pokemon already captured.")
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
		utils.PrintError("Pokemon not captured.")
		return fmt.Errorf("Pokemon not captured")
	}

	delete(pokedex.Pokemons, name)
	return nil
}
