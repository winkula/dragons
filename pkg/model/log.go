// +build !debug

package model

import (
	"log"
	"os"
)

func debug(a ...interface{}) {
	// noop
}

var (
	loge = log.New(os.Stderr, "[error] ", log.Lshortfile)
)
