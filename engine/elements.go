package engine

import pix "github.com/faiface/pixel"

type Dim struct {
	Width, Height float64
}

// Anything that gets updated every frame such as UI elements and Entities etc.
type Component interface {
	Update(*Game) // ? error
}

type HitBoxable interface {
	ToRect() pix.Rect
}

type State int
