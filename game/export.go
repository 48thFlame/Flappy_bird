package game

import "github.com/avitar64/Flappy_bird/engine"

const (
	WindowWidth  = 576
	WindowHeight = 812
	FPS          = 60
)

const (
	GameState engine.State = 1
)

func NewLevel() []engine.Component {
	bgk := newBackground()
	ground := newGround()
	pipeMa := newPipeManager()
	bird := newBird(ground, pipeMa)
	scoreWritter := newScoreWritter()

	return []engine.Component{
		bgk,
		pipeMa,
		ground,
		bird,
		scoreWritter,
	}
}
