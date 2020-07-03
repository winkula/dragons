package main

import (
	"fmt"

	"github.com/winkula/dragons/pkg/model"
)

func main() {
	world := model.NewWorld(3, 3)
	world.SetSquare(0, 0, model.SquareDragon)
	world.SetSquare(0, 1, model.SquareFire)
	world.SetSquare(1, 2, model.SquareEmpty)

	fmt.Printf("World:\n%s\n", world)

	fmt.Printf("Neighbours: %v\n", world.GetNeighbours(0, 0))

	fmt.Printf("Valid: %t\n", model.ValidateWorld(world))

	world2 := model.NewWorld(3, 3)
	successors := world2.Enumerate()
	fmt.Println("Successors:")
	for _, s := range successors {
		fmt.Println(s)
		fmt.Println("----")
	}
}
