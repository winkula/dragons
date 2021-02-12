// +build debug

package model

import (
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
}

func debug(a ...interface{}) {
	log.Println(a...)
}

func debugGrid(g *Grid, i int) {
	debug(Render(g, i))
}

var (
	loge = log.New(os.Stderr, "[error] ", log.Lshortfile)
)
