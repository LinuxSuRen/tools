package main

import (
	"syscall/js"

	"github.com/linuxsuren/tools/pkg"
)

func main() {
	c := make(chan struct{}, 0)

	println("Go WebAssembly Initialized")

	js.Global().Set("formatJSON", js.FuncOf(formatJSON))
	js.Global().Set("base64Encode", js.FuncOf(base64Encode))
	js.Global().Set("base64Decode", js.FuncOf(base64Decode))
	js.Global().Set("generatePassword", js.FuncOf(generatePassword))

	<-c
}

func formatJSON(this js.Value, args []js.Value) interface{} {
	var txt string
	if len(args) > 0 {
		txt, _ = pkg.JsonFormat(args[0].String())
	}
	return js.ValueOf(txt)
}

func base64Encode(this js.Value, args []js.Value) interface{} {
	var txt string
	if len(args) > 0 {
		txt = pkg.Base64Encode(args[0].String())
	}
	return js.ValueOf(txt)
}

func base64Decode(this js.Value, args []js.Value) interface{} {
	var txt string
	if len(args) > 0 {
		txt = pkg.Base64Decode(args[0].String())
	}
	return js.ValueOf(txt)
}

func generatePassword(this js.Value, args []js.Value) interface{} {
	var txt string
	if len(args) >= 3 {
		txt = pkg.GeneratePassword(args[0].Int(), args[1].Bool(), args[2].Bool(), args[3].Bool())
	}
	return js.ValueOf(txt)
}
