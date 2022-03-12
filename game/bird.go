package game

import (
	"github.com/avitar64/Flappy_bird/engine"
	pix "github.com/faiface/pixel"
)

const (
	WindowWidth  = 576
	WindowHeight = 812
	FPS          = 60
)

func NewBird() *engine.Entity {
	e := engine.NewEntity("assets/bird.png", 4)
	e.Pos = pix.V(100, 100)
	e.Expands = append(e.Expands, func(e *engine.Entity) {
		e.Pos.X += 1
	})
	return e
}

func NewBackground() *engine.Entity {
	e := engine.NewEntity("assets/background.png", 4)
	e.Pos = pix.V(WindowWidth/2, WindowHeight/2)
	return e
}
