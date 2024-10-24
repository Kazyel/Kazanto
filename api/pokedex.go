package api

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Kazyel/Poke-CLI/utils"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Pokemon struct {
	Name   string
	Types  []string
	Moves  []string
	Stats  map[string]int
	Height int
	Weight int
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

	tbl := table.New("\nName", "Types", "Height", "Weight")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, pokemon := range pokedex.Pokemons {
		weight := fmt.Sprintf("%.1f kg", float64(pokemon.Weight)/10)
		height := fmt.Sprintf("%.1f m", float64(pokemon.Height)/10)
		types := ""

		for i, pokemonType := range pokemon.Types {
			types += pokemonType

			if i < len(pokemon.Types)-1 {
				types += " & "
			}
		}

		tbl.AddRow(pokemon.Name, types, height, weight)
	}

	tbl.Print()
}

func (pokedex *Pokedex) InspectPokemon(name string) (Pokemon, error) {
	if pokedex.Pokemons[name].Name == "" {
		utils.PrintError("Pokemon not captured yet.")
		return Pokemon{}, fmt.Errorf("Pokemon not captured yet")
	}

	utils.PrintAction("Inspecting "+name, "primary")
	utils.PrintTitle("Pokémon info:")

	color.New(color.FgHiCyan, color.Bold).Fprint(os.Stdout, "Name: ")
	color.New(color.Reset).Fprint(os.Stdout, pokedex.Pokemons[name].Name)

	if len(pokedex.Pokemons[name].Types) > 1 {
		color.New(color.FgHiCyan, color.Bold).Fprint(os.Stdout, "\nTypes: ")
		color.New(color.Reset).Fprint(os.Stdout, pokedex.Pokemons[name].Types[0]+" & "+pokedex.Pokemons[name].Types[1])
	} else {
		color.New(color.FgHiCyan, color.Bold).Fprint(os.Stdout, "\nTypes: ")
		color.New(color.Reset).Fprint(os.Stdout, pokedex.Pokemons[name].Types[0])
	}

	color.New(color.FgHiCyan, color.Bold).Fprint(os.Stdout, "\nStats:\n")
	for stat, value := range pokedex.Pokemons[name].Stats {
		color.New(color.FgHiMagenta, color.Bold).Fprint(os.Stdout, "- "+stat+": ")
		color.New(color.FgHiYellow).Fprint(os.Stdout, strconv.Itoa(value)+"\n")
	}

	return pokedex.Pokemons[name], nil
}

func (pokedex *Pokedex) AddPokemon(pokemon PokemonResponse) error {
	if pokedex.Pokemons[pokemon.Name].Name != "" {
		utils.PrintError("Pokemon already captured.")
		return fmt.Errorf("Pokemon already captured")
	}

	pokemonTypes := make([]string, len(pokemon.Types))
	pokemonMoves := make([]string, len(pokemon.Moves))
	pokemonStats := make(map[string]int, len(pokemon.Stats))

	for i, pokemonType := range pokemon.Types {
		pokemonTypes[i] = pokemonType.Type.Name
	}

	for i, pokemonMove := range pokemon.Moves {
		pokemonMoves[i] = pokemonMove.Move.Name
	}

	for _, pokemonStat := range pokemon.Stats {
		splittedString := strings.Split(pokemonStat.Stat.Name, "-")
		joinedString := strings.Join(splittedString, " ")

		statName := strings.ToUpper(joinedString[0:1]) + joinedString[1:]
		pokemonStats[statName] = pokemonStat.BaseStat
	}

	pokedex.Pokemons[pokemon.Name] = Pokemon{
		Name:   pokemon.Name,
		Types:  pokemonTypes,
		Moves:  pokemonMoves,
		Stats:  pokemonStats,
		Height: pokemon.Height,
		Weight: pokemon.Weight,
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
