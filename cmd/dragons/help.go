package main

import "fmt"

func help() {
	subcommands := []string{
		"parse",
		"generate",
		"enumerate",
		"validate",
		"play",
	}

	fmt.Println(" ┌┬─┐┬─┐┌─┐┌─┐┌─┐┌┐┌┌─┐")
	fmt.Println("  │ │├┬┘├─┤│ ┬│ ││││└─┐")
	fmt.Println(" ─┴─┘┴└─┴ ┴└─┘└─┘┘└┘└─┘")
	fmt.Println("")

	fmt.Println("Valid subcommands are:")
	for _, c := range subcommands {
		fmt.Printf("  %v\n", c)
	}
}
