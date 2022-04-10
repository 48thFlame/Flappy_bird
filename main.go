package main

import (
	"image/color"
	_ "image/png"
	"math/rand"
	"time"

	pix "github.com/faiface/pixel"
	pixgl "github.com/faiface/pixel/pixelgl"

	"github.com/avitar64/Flappy_bird/engine"
	"github.com/avitar64/Flappy_bird/game"
)

func main() {
	pixgl.Run(run)
}

func run() {
	rand.Seed(time.Now().UnixNano())

	winConf := pixgl.WindowConfig{
		Title:  "Flappy Bird",
		Bounds: pix.R(0, 0, game.WindowWidth, game.WindowHeight),
	}
	bgkColor := color.RGBA{
		R: 32,
		G: 32,
		B: 32,
		A: 255,
	}

	g := engine.Initialize(winConf, game.FPS, bgkColor)

	game.MakeGameState(g)
	g.ChangeState(game.GameState)

	game.MakeGameOverState(g)

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
