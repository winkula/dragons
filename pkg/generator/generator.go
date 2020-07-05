package generator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/winkula/dragons/pkg/model"
)

// GenerateWorld generates a random world with the given dimensions.
func GenerateWorld(width int, height int) *model.World {
	// seed the random generator
	rand.Seed(time.Now().UnixNano())

	for {
		world := model.NewWorld(width, height).FillSquares(model.SquareEmpty)
		size := world.Width * world.Height
		numDragons := size / 6
		for i := 0; i < numDragons; {
			index := rand.Intn(size)
			if world.GetSquareByIndex(index) == model.SquareDragon {
				continue
			}
			worldNew := world.Clone()
			worldNew.SetDragon(index)
			if model.ValidateWorld(worldNew) {
				world = worldNew
			} else {
				continue
			}
			i++
		}
		if model.ValidateWorld(world) {
			return world
		}
	}
}

// ObfuscateWorld takes a world and tries to obfuscate it so that some
// information is missing but it still has a definite solution.
func ObfuscateWorld(world *model.World) *model.World {
	size := world.Width * world.Height
	tries := 20
	for i := 0; i < tries; {
		index := rand.Intn(size)
		if world.GetSquareByIndex(index) == model.SquareUndefined {
			continue
		}
		worldNew := world.Clone()
		worldNew.SetSquareByIndex(index, model.SquareUndefined)
		if worldNew.HasDistinctSolution() {
			world = worldNew
			fmt.Println(worldNew)
			i = 0
		}
		i++
	}
	return world
}
