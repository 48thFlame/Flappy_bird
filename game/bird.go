package game

import (
	"fmt"

	"github.com/avitar64/Flappy_bird/engine"
	pix "github.com/faiface/pixel"
	pixgl "github.com/faiface/pixel/pixelgl"
)

const (
	birdGravitySpeed    = .31
	birdMaxGravitySpeed = 9

	// birdRotSpeed    = .46
	// birdMaxRotSpeed = 9
	// birdMaxRot      = 67.5

	scale = 4
)

func newBird(ground *ground) *bird {
	pic, err := engine.LoadPicture("assets/bird.png")
	if err != nil {
		panic(fmt.Errorf("error loading bird sprite: %v", err))
	}

	picRect := pic.Bounds()
	width := picRect.W()
	height := picRect.H()

	return &bird{
		pos: pix.V(WindowWidth/5, WindowHeight/2),
		dim: engine.Dim{
			Width:  width * scale,
			Height: height * scale,
		},
		yv:     0,
		// rot:    0,
		// rv:     0,
		spr:    pix.NewSprite(pic, pic.Bounds()),
		ground: ground,
	}
}

type bird struct {
	spr    *pix.Sprite
	pos    pix.Vec
	yv     float64 // y velocity
	dim    engine.Dim
	rot    float64 // in degrees
	// rv     float64 // rotation velocity
	ground *ground
}

func (b *bird) ToRect() pix.Rect {
	return toRect(b.pos.X, b.pos.Y, b.dim.Width, b.dim.Height)
}

func (b *bird) Update(g *engine.Game) {
	b.movement(g)

	b.spr.Draw(g.Win, pix.IM.Moved(b.pos).Rotated(b.pos, degreesToRadians(b.rot)).Scaled(b.pos, scale))
}

func (b *bird) movement(g *engine.Game) {
	if g.Win.JustReleased(pixgl.MouseButtonLeft) {
		b.yv = birdMaxGravitySpeed
		// b.rv = 0
		// b.rot = birdMaxRot
	}
	// fmt.Println("birdRotChangeSpeed:", birdRotSpeed)

	if b.yv > -birdMaxGravitySpeed {
		b.yv -= birdGravitySpeed
	}

	b.pos.Y += b.yv
	if engine.TouchingEdge(b, g.Win.Bounds().W(), g.Win.Bounds().H()) || engine.Touching(b, b.ground) {
		b.pos.Y -= b.yv
	}

	// if b.rv > -birdMaxRotSpeed {
	// 	b.rv -= birdRotSpeed
	// }

	// b.rot += b.rv
	// if b.rot < -birdMaxRot {
	// 	b.rot -= b.rv
	// }
}
