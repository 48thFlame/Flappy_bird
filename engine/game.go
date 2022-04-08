package engine

import (
	"fmt"
	"image/color"
	"time"

	pixgl "github.com/faiface/pixel/pixelgl"
)

func Initialize(winConf pixgl.WindowConfig, fps int, bgkColor color.RGBA) *Game {
	g := &Game{}

	g.WinConf = winConf

	win, err := pixgl.NewWindow(g.WinConf)
	if err != nil {
		panic(fmt.Errorf("error creating window: %v", err))
	}

	g.Win = win
	g.BgkColor = bgkColor
	g.millsPerFrame = 1000 / float64(fps)
	g.state = 1
	g.states = make(statesType)
	g.statesFields = make(statesFieldsType)

	return g
}

type statesType map[State][]Component
type statesFieldsType map[State]map[string]interface{}

type Game struct {
	WinConf       pixgl.WindowConfig
	Win           *pixgl.Window
	BgkColor      color.RGBA
	millsPerFrame float64
	state         State
	states        statesType
	statesFields  statesFieldsType
}

func (g *Game) update() {
	for _, c := range g.states[g.state] {
		c.Update(g)
	}
}

func (g *Game) Run() {
	for !g.Win.Closed() {
		g.Win.Clear(g.BgkColor)
		g.update()
		g.Win.Update()
		time.Sleep(time.Millisecond * time.Duration(g.millsPerFrame))
	}

	g.Win.Destroy()
}

func (g *Game) AddState(state State) {
	g.states[state] = make([]Component, 0)
}

func (g *Game) AddComponentToState(state State, c ...Component) {
	g.states[state] = append(g.states[state], c...)
}
