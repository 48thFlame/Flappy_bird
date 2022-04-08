package game

import (
	"fmt"

	"github.com/avitar64/Flappy_bird/engine"
	pix "github.com/faiface/pixel"
	pixgl "github.com/faiface/pixel/pixelgl"
)

const (
	birdGravitySpeed    = .45
	birdMaxGravitySpeed = 20

	// birdRotSpeed    = .46
	// birdMaxRotSpeed = 9
	// birdMaxRot      = 67.5

	scale = 4
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
		pos: pix.V(WindowWidth/5, WindowHeight/2),
		dim: engine.Dim{
			Width:  width * scale,
			Height: height * scale,
		},
		yv: 0,
		// rot:    0,
		// rv:     0,
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
	// rv     float64 // rotation velocity
}

func (b *bird) ToRect() pix.Rect {
	return toRect(b.pos.X, b.pos.Y, b.dim.Width, b.dim.Height)
}

func (b *bird) Update(g *engine.Game) {
	b.movement(g)
	// dead := b.pipeCollide()
	// if dead {
	// 	fmt.Println(dead)
	// }

	b.spr.Draw(g.Win, pix.IM.Moved(b.pos).Rotated(b.pos, degreesToRadians(b.rot)).Scaled(b.pos, scale))
}

func (b *bird) movement(g *engine.Game) {
	if g.Win.JustReleased(pixgl.MouseButtonLeft) {
		b.yv = birdMaxGravitySpeed / 2
		// b.rv = 0
		// b.rot = birdMaxRot
	}

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

// func (b *bird) pipeCollide() bool {
// 	for _, p := range b.pipeMa.pipes {
// 		// if toRect(p.bPos.X, p.bPos.Y, p.bDim.Width, p.bDim.Height).Intersects(b.ToRect()) ||
// 		// 	toRect(p.tPos.X, p.tPos.Y, p.tDim.Width, p.tDim.Height).Intersects(b.ToRect()) {
// 		// 	return true
// 		// }
// 		bW, bH := b.dim.Width, b.dim.Height
// 		bR := toRect(p.bPos.X-bW/2, p.bPos.Y-bH/2, bW, bH)

// 		tW, tH := p.tDim.Width, p.tDim.Height
// 		tR := toRect(p.tPos.X-tW/2, p.tPos.Y-tH/2, tW, tH)

// 		if b.ToRect().Intersects(bR) || b.ToRect().Intersects(tR) {
// 			return true
// 		}
// 	}
// 	return false
// }
