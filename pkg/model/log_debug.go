// +build debug

package model

import (
	"log"
	"os"
)

func debug(a ...interface{}) {
	log.Println(a...)
}

var (
	loge = log.New(os.Stderr, "[error] ", log.Lshortfile)
)
