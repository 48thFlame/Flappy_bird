package main

import (
	"fmt"

	pix "github.com/faiface/pixel"
	pixgl "github.com/faiface/pixel/pixelgl"
)

const (
	WindowWidth  = 1024
	WindowHeight = 768
)

func main() {
	pixgl.Run(run)
}

func run() {
	fmt.Println("Hello, World!")

	winConf := pixgl.WindowConfig{
		Title:  "Hello, World!",
		Bounds: pix.R(0, 0, WindowWidth, WindowHeight),
	}

	win, err := pixgl.NewWindow(winConf)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Update()
	}
}
