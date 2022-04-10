package game

import (
	"fmt"
	"math/rand"

	"github.com/avitar64/Flappy_bird/engine"
	pix "github.com/faiface/pixel"
)

const (
	pipeSpeed     = 5.71
	pipeXDistance = 365
	pipeYDistance = 700
)

func newPipePart(name string, x float64) *pipePart {
	pic, err := engine.LoadPicture(fmt.Sprintf("assets/pipe %s.png", name))
	if err != nil {
		panic(fmt.Errorf("error loading pipe sprite: %v", err))
	}

	picRect := pic.Bounds()

	var y float64
	if name == "top" {
		y = pipeYDistance
	} else {
		y = 0
	}

	return &pipePart{
		spr: pix.NewSprite(pic, picRect),
		pos: pix.V(x, y),
		dim: engine.Dim{
			Width:  picRect.W() * scale,
			Height: picRect.H() * scale,
		},
	}
}

type pipePart struct {
	spr *pix.Sprite
	pos pix.Vec
	dim engine.Dim
}

func (pt *pipePart) Update(g *engine.Game) {
	pt.pos.X -= pipeSpeed

	pt.spr.Draw(g.Win, pix.IM.Moved(pt.pos).Scaled(pt.pos, scale))
}

func (pt *pipePart) ToRect() pix.Rect {
	return toRect(pt.pos.X, pt.pos.Y, pt.dim.Width, pt.dim.Height)
}

//

func newPipe(x float64) *pipe {
	top, bottom := newPipePart("top", x), newPipePart("bottom", x)

	return &pipe{
		top:    top,
		bottom: bottom,
	}
}

type pipe struct {
	top, bottom *pipePart
}

func (p *pipe) Update(g *engine.Game) {
	p.top.Update(g)
	p.bottom.Update(g)

	if p.bottom.pos.X < 0 {
		p.bottom.pos.X = WindowWidth + pipeXDistance/2
		p.top.pos.X = p.bottom.pos.X

		setRandYPipe(p)
	}
}

func setRandYPipe(p *pipe) {
	p.bottom.pos.Y = float64(rand.Intn(WindowHeight - 500))
	p.top.pos.Y = p.bottom.pos.Y + pipeYDistance
}

func newPipeManager() *pipeManager {
	pipes := make([]*pipe, 0)

	for i := 0; i < 2; i++ {
		pipe := newPipe(float64(WindowWidth + pipeXDistance*i))
		setRandYPipe(pipe)

		pipes = append(pipes, pipe)
	}

	return &pipeManager{
		pipes: pipes,
	}
}

type pipeManager struct {
	pipes []*pipe
}

func (pm *pipeManager) Update(g *engine.Game) {
	for _, p := range pm.pipes {
		p.Update(g)
	}
}
