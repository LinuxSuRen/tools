package main

import (
	"syscall/js"

	"github.com/linuxsuren/tools/pkg"
)

func main() {
	c := make(chan struct{}, 0)

	println("Go WebAssembly Initialized")

	js.Global().Set("formatJSON", js.FuncOf(formatJSON))

	<-c
}

func formatJSON(this js.Value, args []js.Value) interface{} {
	var txt string
	if len(args) > 0 {
		txt, _ = pkg.JsonFormat(args[0].String())
	}
	return js.ValueOf(txt)
}
