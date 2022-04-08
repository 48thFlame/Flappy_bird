package game

import "github.com/avitar64/Flappy_bird/engine"

const (
	WindowWidth  = 576
	WindowHeight = 812
	FPS          = 60
)

func NewLevel() []engine.Component {
	bgk := newBackground()
	ground := newGround()
	pipeMa := newPipeManager()
	bird := newBird(ground, pipeMa)

	return []engine.Component{
		bgk,
		pipeMa,
		ground,
		bird,
	}
}
