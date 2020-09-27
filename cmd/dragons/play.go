package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/winkula/dragons/pkg/game"
	"github.com/winkula/dragons/pkg/model"
)

func init() {
	cmd := flag.NewFlagSet("play", flag.ExitOnError)

	registerCommand("play", cmd, func() {
		if cmd.NArg() > 0 {
			g := parse(cmd.Arg(0), true)
			playGiven(g)
		} else {
			playNew()
		}
	})
}

func playGiven(g *model.Grid) {
	game := game.NewGameFromPuzzle(g)
	play(game)
}

func playNew() {
	fmt.Println("Loading game...")
	duration := 3 * time.Second
	game := game.NewGame(8, 8, duration)
	play(game)
}

func play(game *game.Game) {
	for {
		clean()

		// render game
		index, _ := game.State.Index(game.CursorX, game.CursorY)
		fmt.Println(model.Render(game.State, nil, index))
		fmt.Println(game)

		// read key
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		switch key {
		case keyboard.KeyArrowLeft:
			game.Left()
		case keyboard.KeyArrowRight:
			game.Right()
		case keyboard.KeyArrowUp:
			game.Up()
		case keyboard.KeyArrowDown:
			game.Down()
		}

		switch char {
		case 'd':
			game.Set(model.SquareDragon)
		case 'f':
			game.Set(model.SquareFire)
		case 'e':
			game.Set(model.SquareEmpty)
		}
	}
}

func clean() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
