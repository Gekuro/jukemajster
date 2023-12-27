package main

import (
	"fmt"
	"image/color"
	"jukemajster/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 800

	CENTRE_POINT_X = WINDOW_WIDTH / 2
	CENTRE_POINT_Y = WINDOW_HEIGHT / 2

	CENTRE_POINT_RADIUS = 50

	GAME_OVER_TEXT = "game over o_o press enter to try again"
)

func main() {
	rl.InitWindow(WINDOW_HEIGHT, WINDOW_WIDTH, "NoName")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	sm := state.NewManager()

	for !rl.WindowShouldClose() {
		if !sm.IsGameOver {
			sm.Update()
		} else {

			if rl.IsKeyDown(rl.KeyEnter) {
				sm = state.NewManager()
			}

		}
		RenderFromState(sm)

	}
}

func RenderFromState(sm *state.Manager) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	defer rl.EndDrawing()

	// player circle
	rl.DrawCircle(sm.PlayerCircle.X, sm.PlayerCircle.Y, sm.PlayerCircle.Radius, sm.PlayerCircle.Color)

	// centre circle
	rl.DrawCircle(CENTRE_POINT_X, CENTRE_POINT_Y, CENTRE_POINT_RADIUS, color.RGBA{R: 50, G: 50, B: 50, A: 255})

	// game over, try again prompt
	if sm.IsGameOver {
		rl.DrawText(GAME_OVER_TEXT, 30, 30, 20, color.RGBA{R: 50, G: 50, B: 50, A: 255})
	}

	// score prompt
	rl.DrawText(fmt.Sprintf("SCORE: %v", sm.Score), 30, 50, 20, color.RGBA{R: 50, G: 50, B: 50, A: 100})

	// bullets
	for _, bullet := range sm.Bullets {
		rl.DrawCircle(bullet.X, bullet.Y, bullet.ModelRadius, bullet.Color)
	}
}
