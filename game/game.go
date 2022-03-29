package game

import (
	"github.com/avitar64/Flappy_bird/engine"
	pixgl "github.com/faiface/pixel/pixelgl"
)

const (
	WindowWidth  = 576
	WindowHeight = 812
	FPS          = 60
)

const (
	scale        = 4
)

func NewGame(win *pixgl.Window) []engine.Component {
	entSlice := []*engine.Entity{
		newBackground(),
		newGround(),
		newBird(win),
	}

	// this part need because []struct != []interface{} that it implements, see:
	// https://stackoverflow.com/questions/12994679/slice-of-struct-slice-of-interface-it-implements
	compSlice := make([]engine.Component, 0)

	for _, ent := range entSlice {
		compSlice = append(compSlice, engine.Component(ent))
	}

	return compSlice
}
