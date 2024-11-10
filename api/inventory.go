package api

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Inventory struct {
	Pokeballs map[string]*Pokeballs
}

type Pokeballs struct {
	Pokeball Pokeball
	Quantity int
}

func (player *Player) RenderInventory() {
	headerFmt := color.New(color.FgHiBlue, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	pokeballTable := table.New("Type", "Quantity")
	pokeballTable.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, pokeball := range player.Inventory.Pokeballs {
		pokeballTable.AddRow(pokeball.Pokeball.Name, pokeball.Quantity)
	}

	pokeballTable.Print()
}
