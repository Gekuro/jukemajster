package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"noname/models"
	"image/color"
)

const(
	WINDOW_WIDTH = 800
	WINDOW_HEIGHT = 800

	CENTRE_POINT_X = WINDOW_WIDTH / 2
	CENTRE_POINT_Y = WINDOW_HEIGHT / 2

	PLAYER_GAME_RADIUS = 300

	PLAYER_MOVING_SPEED = 0.03

	CENTRE_POINT_RADIUS = 30

) 

func main() {
	rl.InitWindow(WINDOW_HEIGHT, WINDOW_WIDTH, "NoName")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	pc := models.NewPlayerCircle(10, CENTRE_POINT_X, CENTRE_POINT_Y, PLAYER_GAME_RADIUS)

	bullets := make([]*models.Bullet, 0)
	var newBulletCounter int

	for !rl.WindowShouldClose() {
		// update ----------------------------------------------------------------
		// get input and update the circle pos using very cool equation
		pc.UpdatePos();

		if rl.IsKeyDown(rl.KeyUp) {
			pc.Angle += PLAYER_MOVING_SPEED
		}
		if rl.IsKeyDown(rl.KeyDown) {
			pc.Angle -= PLAYER_MOVING_SPEED
		}

		// every n-th frame create a new bullet
		newBulletCounter += 1
		if newBulletCounter > models.FRAMES_BETWEEN_BULLETS {
			newBulletCounter = 0
			bullets = append(bullets, models.NewBullet(models.BULLET_RADIUS, models.BULLET_SPEED, pc.Angle, CENTRE_POINT_X, CENTRE_POINT_Y))
		}
	
		// render ----------------------------------------------------------------

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// draws the point of the player circle
		rl.DrawCircle(pc.X, pc.Y, pc.Radius, pc.Color)


		// draws the centre circle
		rl.DrawCircle(CENTRE_POINT_X, CENTRE_POINT_Y, CENTRE_POINT_RADIUS, color.RGBA{R: 50, G: 50, B: 50, A: 255})

		// moves and draws all existing bullets
		for _, bullet := range bullets {
			bullet.UpdatePos()
			// check if bullet radius is over some value, then delete it
			rl.DrawCircle(bullet.X, bullet.Y, bullet.ModelRadius, bullet.Color)
		}

		rl.EndDrawing()
	}
}