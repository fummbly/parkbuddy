package main

import (
	"syscall/js"
)

func main() {
	doc := js.Global().Get("document")

	canvasEl := doc.Call("getElementById", "Canvas")
	width := doc.Get("body").Get("clientWidth").Float()
	height := doc.Get("body").Get("clientHeight").Float()

	canvasEl.Set("width", width)
	canvasEl.Set("height", height)
	ctx := canvasEl.Call("getContext", "2d")

	ctx.Set("fillStyle", "green")
	ctx.Call("fillRect", 10, 10, 150, 100)

}
