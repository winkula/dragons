package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	downloadWasmExec()
	fileServer := http.FileServer(http.Dir("./web"))
	http.Handle("/", fileServer)
	fmt.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", nil)
}

const wasmExecURL = "https://raw.githubusercontent.com/golang/go/release-branch.go1.15/misc/wasm/wasm_exec.js"
const wasmExecFile = "web/wasm_exec.js"

func downloadWasmExec() {
	if _, err := os.Stat(wasmExecFile); err == nil {
		return
	}
	println("Downloading wasm_exec.js...")
	out, err := os.Create(wasmExecFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	resp, err := http.Get(wasmExecURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
