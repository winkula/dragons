// +build !debug

package model

import (
	"log"
	"os"
)

type dummyWriter struct{}

func (w *dummyWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

var (
	logd = log.New(&dummyWriter{}, "[debug] ", log.Lshortfile)
	loge = log.New(os.Stderr, "[error] ", log.Lshortfile)
)
