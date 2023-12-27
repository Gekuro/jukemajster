package models

import (
	"image/color"
	"math"
	"math/rand"
)

const (
	FRAMES_BETWEEN_BULLETS = 5
	BULLET_SPEED           = 7
	BULLET_RADIUS          = 7
	BULLET_ACCURACY        = 2 * math.Pi / 2
)

type Bullet struct {
	X              int32
	Y              int32
	ModelRadius    float32
	centerX        int32
	centerY        int32
	PositionRadius float64
	Angle          float64
	Speed          float64
	Color          color.RGBA
}

func (b *Bullet) UpdatePos() {
	b.PositionRadius += b.Speed

	b.X = b.centerX + int32(b.PositionRadius*math.Sin(b.Angle))
	b.Y = b.centerY + int32(b.PositionRadius*math.Cos(b.Angle))
}

func NewBullet(br float32, bs float64, targetAngle float64, cx int32, cy int32) *Bullet {
	randomFloat := rand.Float64()*(BULLET_ACCURACY) - (BULLET_ACCURACY / 2)
	randomAngle := targetAngle + randomFloat

	b := &Bullet{
		X:              0,
		Y:              0,
		ModelRadius:    br,
		centerX:        cx,
		centerY:        cy,
		PositionRadius: 0,
		Angle:          randomAngle,
		Speed:          bs,
		Color:          color.RGBA{R: 0, G: 0, B: 0, A: 255},
	}
	b.UpdatePos()
	return b
}
