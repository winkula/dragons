package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/winkula/dragons/pkg/game"
	"github.com/winkula/dragons/pkg/model"
)

func play() {
	fmt.Println("Loading game...")

	duration := 1 * time.Second
	game := game.NewGame(8, 8, duration)

	for {
		// clean
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()

		// render game
		index := game.State.Height*game.CursorY + game.CursorX
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
