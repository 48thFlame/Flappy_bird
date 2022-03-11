package engine

import (
	"fmt"

	pix "github.com/faiface/pixel"
)

type Expansion func(*Entity)

type ScaleType struct {
	X, Y float64
}

func NewEntity(path string) *Entity {
	e := &Entity{}

	pic, err := loadPicture(path)
	if err != nil {
		panic(fmt.Errorf("error loading sprite: %v", err))
	}

	sprite := pix.NewSprite(pic, pic.Bounds())

	e.Spr = sprite
	e.Pos = pix.V(0, 0)
	e.Rot = 0
	e.Pic = pic
	e.Scale = ScaleType{X: 0, Y: 0}
	e.Expands = make([]Expansion, 0)
	e.Fields = make(map[string]interface{})

	return e
}

type Entity struct {
	Spr     *pix.Sprite // sprite
	Pos     pix.Vec     // postition
	Rot     float64     // rotation
	Pic     pix.Picture // image
	Scale   ScaleType   // dimensions
	Expands []Expansion
	Fields  map[string]interface{}
}

func (e *Entity) Update(g *Game) {
	for _, expand := range e.Expands {
		expand(e)
	}
	e.Spr.Draw(g.Win, pix.IM.Moved(e.Pos).Rotated(e.Pos, e.Rot).ScaledXY(e.Pos, pix.V(e.Scale.X, e.Scale.Y)))
}

func EntityCollides(e1, e2 *Entity) bool {
	rect1 := e1.Pic.Bounds()
	rect2 := e2.Pic.Bounds()

	return rect1.Intersects(rect2)
}
