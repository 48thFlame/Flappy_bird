package game

import (
	"github.com/avitar64/Flappy_bird/engine"
	pix "github.com/faiface/pixel"
	pixgl "github.com/faiface/pixel/pixelgl"
)

const (
	birdVelName         = "vel"
	birdWinName         = "win"
	birdGravitySpeed    = .315
	birdGravityMaxSpeed = 8.76
)

// func entTouchingEdge(e *engine.Entity) bool {

// a function that return weather the entity is touching any edge of the window taken in count the width and height of the entity
func entTouchingEdge(e *engine.Entity) bool {
	return (e.Pos.X < e.Dim.Width*e.Scale/2 ||
		e.Pos.X > WindowWidth-e.Dim.Width*e.Scale/2 ||
		e.Pos.Y < e.Dim.Height*e.Scale/2 ||
		e.Pos.Y > WindowHeight-e.Dim.Height*e.Scale/2)
}

func newBird(win *pixgl.Window) *engine.Entity {
	bird := engine.NewEntity("assets/bird.png", 4)
	bird.Pos = pix.V(bird.Dim.Width*bird.Scale*2, WindowHeight/2+bird.Dim.Height*bird.Scale)
	// bird.Pos = pix.V(0, 0)
	// fmt.Println("touching edge: ", entTouchingEdge(bird))

	initBirdFields(bird, win)

	bird.Expands = append(bird.Expands, birdMovement)

	return bird
}

func initBirdFields(e *engine.Entity, win *pixgl.Window) {
	fields := e.Fields

	fields[birdVelName] = 0.0
	fields[birdWinName] = win
}

func birdMovement(e *engine.Entity) {
	win := e.Fields[birdWinName].(*pixgl.Window)
	vel := e.Fields[birdVelName].(float64)

	if vel > -birdGravityMaxSpeed {
		vel += birdGravitySpeed
	}

	if !entTouchingEdge(e) {
		e.Pos.Y -= vel

		if win.JustPressed(pixgl.KeySpace) {
			vel = birdGravityMaxSpeed
		}
	}

	e.Fields[birdVelName] = vel
	
	// if vel > -birdGravityMaxSpeed { // if velocity is more than -gravityMaxSpeed, then speed up the fall
	// 	vel -= birdGravitySpeed
	// }

	
	// // if entTouchingEdge(e) {
	// e.Pos.Y += vel
	// // e.Pos.Y -= vel
	// vel += birdGravitySpeed

	// if win.JustReleased(pixgl.KeySpace) {
	// 	vel = birdGravityMaxSpeed
	// }
	// // }

	// // if e.Pos.Y > WindowHeight-e.Dim.Height*e.Scale/2 || e.Pos.Y < groundHeight {
	// // 	e.Pos.Y -= vel
	// // } else {
	// // 	if win.JustReleased(pixgl.MouseButtonLeft) {
	// // 		vel = birdGravityMaxSpeed
	// // 	}
	// // }
}
