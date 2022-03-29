package game

import (
	"github.com/avitar64/Flappy_bird/engine"
	pix "github.com/faiface/pixel"
)

const (
	fieldsStartPosName = "startPos"

	groundMovementSpeed = 4
	groundHeight = 16
)

func newBackground() *engine.Entity {
	e := engine.NewEntity("assets/background.png", scale)
	e.Pos = pix.V(WindowWidth/2, WindowHeight/2)

	return e
}

func newGround() *engine.Entity {
	e := engine.NewEntity("assets/ground.png", scale)

	startPos := pix.V(WindowWidth-e.Dim.Width*e.Scale/2+40, groundHeight)

	e.Fields[fieldsStartPosName] = startPos
	e.Pos = startPos

	e.Expands = append(e.Expands, groundMovement)

	return e
}

func groundMovement(e *engine.Entity) {
	if e.Pos.X < WindowWidth/2 {
		e.Pos = e.Fields[fieldsStartPosName].(pix.Vec)
	}
	
	e.Pos.X -= groundMovementSpeed
}
