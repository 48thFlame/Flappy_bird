package game

import (
	"fmt"

	"github.com/avitar64/Flappy_bird/engine"
	pix "github.com/faiface/pixel"
)

const (
	birdGravitySpeed    = .65
	birdMaxGravitySpeed = 20
	birdJumpSpeed       = birdMaxGravitySpeed / 1.75
)

func newBird(ground *ground, pipeMa *pipeManager) *bird {
	pic, err := engine.LoadPicture("assets/bird.png")
	if err != nil {
		panic(fmt.Errorf("error loading bird sprite: %v", err))
	}

	picRect := pic.Bounds()
	width := picRect.W()
	height := picRect.H()

	return &bird{
		pos: pix.V(WindowWidth/5, WindowHeight-128),
		dim: engine.Dim{
			Width:  width * scale,
			Height: height * scale,
		},
		yv:     0,
		spr:    pix.NewSprite(pic, pic.Bounds()),
		ground: ground,
		pipeMa: pipeMa,
	}
}

type bird struct {
	spr    *pix.Sprite
	pos    pix.Vec
	yv     float64 // y velocity
	dim    engine.Dim
	rot    float64 // in degrees
	ground *ground
	pipeMa *pipeManager
}

func (b *bird) ToRect() pix.Rect {
	return toRect(b.pos.X, b.pos.Y, b.dim.Width, b.dim.Height)
}

func (b *bird) Update(g *engine.Game) {
	b.movement(g)

	gameOver := g.GetStateField(GameState, "gameOver").(bool)
	yDifferBG := b.pos.Y - b.ground.back.pos.Y
	if gameOver && yDifferBG < b.dim.Height+b.ground.dim.Height/2 { // if game ended and bird fell the hole way down
		g.ChangeState(GameOver)
	}

	dead := b.pipeCollide()
	if dead {
		g.SetStateField(GameState, "gameOver", true)
	}

	b.incremtnScore(g)

	b.spr.Draw(g.Win, pix.IM.Moved(b.pos).Rotated(b.pos, degreesToRadians(b.rot)).Scaled(b.pos, scale))
}

func (b *bird) movement(g *engine.Game) {
	dead := g.GetStateField(GameState, "gameOver").(bool)

	if !dead {
		if gotUserInput(g.Win) {
			b.yv = birdJumpSpeed
		}
	}

	if b.yv > -birdMaxGravitySpeed {
		b.yv -= birdGravitySpeed
	}

	b.pos.Y += b.yv
	if engine.TouchingEdge(b, g.Win.Bounds().W(), g.Win.Bounds().H()) || engine.Touching(b, b.ground) {
		b.pos.Y -= b.yv
	}
}

func (b *bird) pipeCollide() bool {
	for _, pipe := range b.pipeMa.pipes {
		if engine.Touching(b, pipe.bottom) || engine.Touching(b, pipe.top) {
			return true
		}
	}

	return false
}

func (b *bird) incremtnScore(g *engine.Game) {
	for _, pipe := range b.pipeMa.pipes {
		posDifference := b.pos.X - pipe.bottom.pos.X
		// cant just check if bird.pos.x is bigger then pipe, so only adds score for each pipe once
		if posDifference > 0 && posDifference < pipeSpeed {
			score := g.GetStateField(GameState, "score").(int)
			g.SetStateField(GameState, "score", score+1)
		}
	}
}
