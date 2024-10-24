package utils

import (
	"os"

	"github.com/fatih/color"
)

func PrintError(input string) {
	color.New(color.FgRed, color.Bold).Fprint(os.Stdout, "\n[!] ")
	color.New(color.FgRed).Fprintln(os.Stdout, input)
}

func PrintCommmands(name string, description string) {
	color.New(color.FgHiMagenta, color.Bold).Fprint(os.Stdout, name+": ")
	color.New(color.FgHiYellow).Fprint(os.Stdout, description+"\n")
}

func PrintTitle(input string) {
	formatedInput := "\n" + input + "\n"
	color.New(color.FgHiRed, color.Bold).Fprintf(os.Stdout, formatedInput)
}

func PrintSuccessfulCatch() {
	color.New(color.FgHiYellow, color.Bold).Fprintln(os.Stdout, "You caught the Pokemon!")
}

func PrintFailedCatch(pokemonName string) {
	color.New(color.FgHiRed).Fprintf(os.Stdout, "\nYou failed to catch the %s.\n", pokemonName)
}

func PrintAction(action string, isSecondary string) {
	formattedAction := action + "...\n"

	if isSecondary == "secondary" {
		color.New(color.FgHiGreen, color.Bold).Fprintf(os.Stdout, formattedAction)
		return
	}

	color.New(color.FgGreen).Fprintf(os.Stdout, formattedAction)
}

func PrintCachedAction(action string) {
	formattedAction := action + "...\n"

	color.New(color.FgHiGreen, color.Bold).Fprintf(os.Stdout, "[CACHED] ")
	color.New(color.FgGreen).Fprintf(os.Stdout, formattedAction)
}
