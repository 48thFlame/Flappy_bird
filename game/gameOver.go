package game

import (
	"fmt"

	"github.com/avitar64/Flappy_bird/engine"
	pix "github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

var (
	gameOverTextLines = []string{
		"Game Over!",
		"Click, or press ",
		"space to restart.",
		"",
		"",
	}
)

const (
	scoreGameOverLineI = 4
)

func NewGameOver() *gameOver {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	text := text.New(pix.V(WindowWidth/2, WindowHeight/2+128), atlas)
	pos := text.Orig

	return &gameOver{
		pos:  pos,
		text: text,
	}
}

type gameOver struct {
	text *text.Text
	pos  pix.Vec
}

func (gov *gameOver) Update(g *engine.Game) {
	gov.text.Clear()
	for i, line := range gameOverTextLines {

		lineToWrite := line

		if i == scoreGameOverLineI {
			lineToWrite = fmt.Sprintf("You scored: %d", g.GetStateField(GameState, "score").(int))
		}

		gov.text.Dot.X -= gov.text.BoundsOf(lineToWrite).W() / 2
		fmt.Fprintln(gov.text, lineToWrite)
	}

	gov.text.Draw(g.Win, pix.IM.Scaled(gov.pos, scale))

	if gotUserInput(g.Win) {
		MakeGameState(g)
		g.ChangeState(GameState)
	}
}
