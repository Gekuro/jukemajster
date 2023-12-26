package models

import (
	"image/color"
	"math"
)

const (
	PLAYER_RGB_R = 255
	PLAYER_RGB_G = 0
	PLAYER_RGB_B = 0
)

type PlayerCircle struct {
	X int32
	Y int32

	centerX int32
	centerY int32
	gameRadius float64
	
	Radius float32
	Angle float64
	Color color.RGBA
}

func (c *PlayerCircle) UpdatePos() {
	c.X = c.centerX + int32(c.gameRadius * math.Sin(c.Angle))
	c.Y = c.centerY + int32(c.gameRadius * math.Cos(c.Angle))
}

func NewPlayerCircle(rad float32, cx int32, cy int32, gr float64) *PlayerCircle {
	pc := &PlayerCircle{
		X: 0, 
		Y: 0, 
		
		centerX: cx,
		centerY: cy,
		gameRadius: gr,

		Radius: rad, 
		Angle: 0, 
		Color: color.RGBA{R: PLAYER_RGB_R, G: PLAYER_RGB_G, B: PLAYER_RGB_B, A: 255},
	}
	pc.UpdatePos()
	return pc
}