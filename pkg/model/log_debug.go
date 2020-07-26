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

var (
	loge = log.New(os.Stderr, "[error] ", log.Lshortfile)
)
