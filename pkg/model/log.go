// +build !debug

package model

import (
	"log"
	"os"
)

func debug(a ...interface{}) {
	// noop
}

func debugGrid(g *Grid, i int) {
	// noop
}

var (
	loge = log.New(os.Stderr, "[error] ", log.Lshortfile)
)
