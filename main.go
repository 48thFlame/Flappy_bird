package main

import (
	"image/color"
	_ "image/png"

	pix "github.com/faiface/pixel"
	pixgl "github.com/faiface/pixel/pixelgl"

	"github.com/avitar64/Flappy_bird/engine"
)

const (
	WindowWidth  = 624
	WindowHeight = 916
	FPS          = 60
)

func main() {
	pixgl.Run(run)
}

const (
	game engine.State = 1
)

func run() {
	winConf := pixgl.WindowConfig{
		Title:  "Hello, World!",
		Bounds: pix.R(0, 0, WindowWidth, WindowHeight),
	}
	bgkColor := color.RGBA{
		R: 32,
		G: 32,
		B: 32,
		A: 255,
	}

	g := engine.Initialize(winConf, FPS, bgkColor)
	g.AddState(game)
	g.ChangeState(game)
	g.Run()
}

// multiple windows:

// for i := 0; i < 5; i++ {
// 	go func() {
// 		g := engine.Initialize(winConf, FPS, bgkColor)
// 		g.Run()
// 	}()
// }
// fmt.Scanln()
