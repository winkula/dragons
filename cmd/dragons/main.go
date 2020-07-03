package main

import (
	"fmt"

	"github.com/winkula/dragons/pkg/model"
)

func main() {
	world := model.NewWorld(3, 3)
	world.SetField(0, 0, 3)
	world.SetField(0, 1, 2)
	world.SetField(1, 2, 1)

	fmt.Printf("World:\n%s\n", world)

	fmt.Printf("Neighbours: %v\n", world.GetNeighbours(0, 0))

	fmt.Printf("Valid: %t\n", model.ValidateWorld(world))
}
