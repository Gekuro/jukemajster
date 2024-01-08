package state

import (
	"jukemajster/models"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CIRCLE_HIT_ANGLE = 2 * math.Pi / 2000

	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 800

	CENTRE_POINT_X = WINDOW_WIDTH / 2
	CENTRE_POINT_Y = WINDOW_HEIGHT / 2

	PLAYER_GAME_RADIUS  = 300
	BULLET_DEATH_RADIUS = 800

	PLAYER_MOVING_SPEED = 0.03
)

type Manager struct {
	PlayerCircle     *models.PlayerCircle
	Bullets          []*models.Bullet
	newBulletCounter int
	Score            int
	IsGameOver       bool
}

func (m *Manager) Update() {
	// check input, and update player position
	if rl.IsKeyDown(rl.KeyUp) {
		m.PlayerCircle.Angle += PLAYER_MOVING_SPEED
	}
	if rl.IsKeyDown(rl.KeyDown) {
		m.PlayerCircle.Angle -= PLAYER_MOVING_SPEED
	}
	m.PlayerCircle.UpdatePos()

	// remove out-of-range bullets
	remainingBullets := make([]*models.Bullet, 0)
	for _, bullet := range m.Bullets {
		bullet.UpdatePos()

		if bullet.PositionRadius < BULLET_DEATH_RADIUS {
			remainingBullets = append(remainingBullets, bullet)
		}
	}
	m.Bullets = remainingBullets

	// every n-th frame create a new bullet
	m.newBulletCounter += 1
	if m.newBulletCounter > models.FRAMES_BETWEEN_BULLETS {
		m.newBulletCounter = 0

		m.Score += 1

		m.Bullets = append(
			m.Bullets,
			models.NewBullet(
				models.BULLET_RADIUS,
				models.BULLET_SPEED,
				m.PlayerCircle.Angle,
				CENTRE_POINT_X,
				CENTRE_POINT_Y,
			),
		)
	}

	// collision score counter
	for _, bullet := range m.Bullets {
		if m.IsGameOver {
			continue
		}

		// TODO clean this monstrosity
		bulletIntersectsPlayerXLowerBound := (float32(bullet.X) > (float32(m.PlayerCircle.X) - (bullet.ModelRadius + m.PlayerCircle.Radius)))
		bulletIntersectsPlayerXUpperBound := (float32(bullet.X) < (float32(m.PlayerCircle.X) + (bullet.ModelRadius + m.PlayerCircle.Radius)))
		bulletIntersectsPlayerYLowerBound := (float32(bullet.Y) > (float32(m.PlayerCircle.Y) - (bullet.ModelRadius + m.PlayerCircle.Radius)))
		bulletIntersectsPlayerYUpperBound := (float32(bullet.Y) < (float32(m.PlayerCircle.Y) + (bullet.ModelRadius + m.PlayerCircle.Radius)))
		bulletIntersectsPlayerX := bulletIntersectsPlayerXLowerBound && bulletIntersectsPlayerXUpperBound
		bulletIntersectsPlayerY := bulletIntersectsPlayerYLowerBound && bulletIntersectsPlayerYUpperBound

		if bulletIntersectsPlayerX && bulletIntersectsPlayerY {
			m.IsGameOver = true
		}
	}
}

func NewManager() *Manager {
	pc := models.NewPlayerCircle(10, CENTRE_POINT_X, CENTRE_POINT_Y, PLAYER_GAME_RADIUS)

	bullets := make([]*models.Bullet, 0)

	return &Manager{
		PlayerCircle:     pc,
		Bullets:          bullets,
		newBulletCounter: 0,
		Score:            0,
		IsGameOver:       false,
	}
}
