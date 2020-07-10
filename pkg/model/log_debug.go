// +build debug

package model

import (
	"log"
	"os"
)

var (
	logd = log.New(os.Stderr, "[debug] ", log.Lshortfile)
	loge = log.New(os.Stderr, "[error] ", log.Lshortfile)
)
