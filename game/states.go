package game

import "github.com/avitar64/Flappy_bird/engine"

const (
	GameState engine.State = 1
	GameOver  engine.State = 2
)

func MakeGameState(g *engine.Game) {
	g.CreateState(GameState)
	g.AddComponentsToState(GameState, NewLevel()...)
	g.SetStateField(GameState, "score", 0)
	g.SetStateField(GameState, "gameOver", false)
}

func MakeGameOverState(g *engine.Game) {
	g.CreateState(GameOver)
	g.AddComponentsToState(GameOver, NewGameOver())
}
