package game

import (
	"fmt"

	"github.com/avitar64/Flappy_bird/engine"

	pix "github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

var (
	digits        = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	scoreTextSpot = pix.V(20, WindowHeight-60)
)

func newScoreWritter() *scoreWritter {
	atlas := text.NewAtlas(basicfont.Face7x13, digits)
	text := text.New(scoreTextSpot, atlas)
	pos := text.Orig

	return &scoreWritter{
		pos:  pos,
		text: text,
	}
}

type scoreWritter struct {
	pos  pix.Vec
	text *text.Text
}

func (sw *scoreWritter) Update(g *engine.Game) {
	score := g.GetStateField(GameState, "score").(int)

	sw.text.Clear()
	fmt.Fprint(sw.text, score)

	sw.text.Draw(g.Win, pix.IM.Scaled(sw.pos, scale))
}
