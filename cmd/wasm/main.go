package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/winkula/dragons/pkg/model"
)

func add(this js.Value, i []js.Value) interface{} {
	fmt.Println("add was called!")
	fmt.Println(js.ValueOf(i[0].Int() + i[1].Int()).String())
	return i[0].Int() + i[1].Int()
}

func parse(this js.Value, i []js.Value) interface{} {
	g := model.Parse(i[0].String())
	j, _ := json.Marshal(g)
	return string(j)
}

func main() {
	quit := make(chan struct{}, 0)

	fmt.Println("Running WASM..")

	fmt.Println("Registering callbacks...")
	js.Global().Get("dragons").Set("add", js.FuncOf(add))
	js.Global().Get("dragons").Set("parse", js.FuncOf(parse))
	js.Global().Get("dragons").Set("stop", js.FuncOf(func(js.Value, []js.Value) interface{} {
		quit <- struct{}{}
		return nil
	}))

	<-quit
}
