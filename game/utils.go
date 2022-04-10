package game

import (
	"math"

	pix "github.com/faiface/pixel"
	pixgl "github.com/faiface/pixel/pixelgl"
)

const (
	scale = 4
)

func toRect(x, y, w, h float64) pix.Rect {
	hw, hh := w/2, h/2

	return pix.R(x-hw, y-hh, x+hw, y+hh)
}

func degreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

func gotUserInput(win *pixgl.Window) bool {
	return win.JustReleased(pixgl.KeySpace) || win.JustReleased(pixgl.MouseButtonLeft)
}
