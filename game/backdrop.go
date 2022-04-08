package game

import (
	"fmt"

	"github.com/avitar64/Flappy_bird/engine"
	pix "github.com/faiface/pixel"
)

type backdrop struct {
	spr *pix.Sprite
	pos pix.Vec
}

func (b *backdrop) Update(g *engine.Game) {
	b.spr.Draw(g.Win, pix.IM.Moved(b.pos).Scaled(b.pos, scale))
}

func newBackground() *backdrop {
	pic, err := engine.LoadPicture("assets/background.png")
	if err != nil {
		panic(fmt.Errorf("error loading background sprite: %v", err))
	}

	return &backdrop{
		spr: pix.NewSprite(pic, pic.Bounds()),
		pos: pix.V(WindowWidth/2, WindowHeight/2),
	}
}

type ground struct {
	back *backdrop
	dim  *engine.Dim
}

func (gr *ground) Update(g *engine.Game) {
	gr.back.Update(g)
}

func (gr *ground) ToRect() pix.Rect {
	return toRect(gr.back.pos.X, gr.back.pos.Y, gr.dim.Width, gr.dim.Height)
}

func newGround() *ground {
	pic, err := engine.LoadPicture("assets/ground.png")
	if err != nil {
		panic(fmt.Errorf("error loading ground sprite: %v", err))
	}

	picRect := pic.Bounds()
	width := picRect.W()
	height := picRect.H()

	return &ground{
		back: &backdrop{
			spr: pix.NewSprite(pic, pic.Bounds()),
			pos: pix.V(WindowWidth/2, WindowHeight/42),
		},
		dim: &engine.Dim{
			Width:  width * scale,
			Height: height * scale,
		},
	}
}
