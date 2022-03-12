package engine

import (
	"fmt"

	pix "github.com/faiface/pixel"
)

type Expansion func(*Entity)

type Dimension struct {
	Width, Height float64
}

func NewEntity(path string, scale float64) *Entity {
	e := &Entity{}

	pic, err := loadPicture(path)
	if err != nil {
		panic(fmt.Errorf("error loading sprite: %v", err))
	}

	picRect := pic.Bounds()
	dim := Dimension{}
	dim.Width = picRect.W()
	dim.Height = picRect.H()

	sprite := pix.NewSprite(pic, pic.Bounds())

	e.Spr = sprite
	e.Pos = pix.V(0, 0)
	e.Rot = 0
	e.Pic = pic
	e.Scale = scale
	e.Dim = dim
	e.Expands = make([]Expansion, 0)
	e.Fields = make(map[string]interface{})

	return e
}

type Entity struct {
	Spr     *pix.Sprite // sprite
	Pos     pix.Vec     // postition
	Rot     float64     // rotation
	Pic     pix.Picture // image
	Scale   float64	 // scale
	Dim     Dimension // dimensions
	Expands []Expansion
	Fields  map[string]interface{}
}

func (e *Entity) Update(g *Game) {
	for _, expand := range e.Expands {
		expand(e)
	}
	e.Spr.Draw(g.Win, pix.IM.Moved(e.Pos).Rotated(e.Pos, e.Rot).Scaled(e.Pos, e.Scale)) // ? pix.ZV in Scaled?
}

func EntityCollides(e1, e2 *Entity) bool {
	rect1 := entToRect(e1)
	rect2 := entToRect(e2)

	return rect1.Intersects(rect2)
}

func entToRect(e *Entity) pix.Rect {
	return pix.R(e.Pos.X, e.Pos.Y, e.Pos.X+e.Dim.Width, e.Pos.Y+e.Dim.Height)
}
