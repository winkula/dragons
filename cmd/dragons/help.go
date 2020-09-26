package main

import "fmt"

func help() {
	fmt.Println(" ┌┬─┐┬─┐┌─┐┌─┐┌─┐┌┐┌┌─┐")
	fmt.Println("  │ │├┬┘├─┤│ ┬│ ││││└─┐")
	fmt.Println(" ─┴─┘┴└─┴ ┴└─┘└─┘┘└┘└─┘")
	fmt.Println("")

	fmt.Println("Valid subcommands are:")
	for name, _ := range commands {
		fmt.Printf("  %v\n", name)
	}
}
