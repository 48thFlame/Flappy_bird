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
	g.StatesComponents = make(StateComponentsType)
	g.StatesFields = make(StatesFieldsType)

	return g
}

type StateComponentsType map[State][]Component
type StatesFieldsType map[State]map[string]interface{}

type Game struct {
	WinConf          pixgl.WindowConfig
	Win              *pixgl.Window
	BgkColor         color.RGBA
	millsPerFrame    float64
	state            State
	StatesComponents StateComponentsType
	StatesFields     StatesFieldsType
}

func (g *Game) update() {
	g.Win.Clear(g.BgkColor)

	for _, component := range g.StatesComponents[g.state] {
		component.Update(g)
	}

	g.Win.Update()
}

func (g *Game) AddState(state State) {
	g.StatesComponents[state] = make([]Component, 0)
}

func (g *Game) AddComponentToState(state State, components ...Component) {
	g.StatesComponents[state] = append(g.StatesComponents[state], components...)
}

func (g *Game) ChangeState(state State) {
	g.state = state
}

func (g *Game) Run() {
	for !g.Win.Closed() {
		g.update()
		time.Sleep(time.Millisecond * time.Duration(g.millsPerFrame))
	}
	g.Win.Destroy()
}
