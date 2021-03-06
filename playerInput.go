package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

const controlerStartAngle, controlerStartPower float64 = 0, 1
const defaultAngleLerp, defaultPowerLerp float64 = math.Pi / 96, 0.1
const defaultMaxPower float64 = 20
const defaultIndicatorDistanceFromBall float64 = 10

const keyFullscreenToggle=ebiten.KeyF
const keyMainMenu=ebiten.Key1
const keyNextTrack=ebiten.KeyBackslash

var mouseBtnIsDown=false




type controler struct {
	angle float64
	power float64

	_angleLerp float64
	_powerLerp float64
	_maxPower  float64

	powerUpKey, powerDownKey ebiten.Key
	angleUpKey, angleDownKey ebiten.Key
	hitKey                   ebiten.Key
	indicator                *angleIndicator

	keyDownMap map[ebiten.Key]bool

	//makes fentching position of ball cleaner in some places
	parent *ball
}

type angleIndicator struct {
	position         vector2
	graphic          *ebiten.Image
	opts             *ebiten.DrawImageOptions
	distanceFromBall float64
}

func makeControler(b *ball) *controler {
	c := new(controler)
	c.angle = controlerStartAngle
	c.power = controlerStartPower
	c._maxPower = defaultMaxPower
	c._angleLerp = defaultAngleLerp
	c._powerLerp = defaultPowerLerp
	c.powerUpKey = ebiten.KeyUp
	c.powerDownKey = ebiten.KeyDown
	c.angleUpKey = ebiten.KeyLeft
	c.angleDownKey = ebiten.KeyRight
	c.hitKey = ebiten.KeySpace
	c.parent = b

	c.keyDownMap=make(map[ebiten.Key]bool)
	c.keyDownMap[keyFullscreenToggle]=false
	c.keyDownMap[keyMainMenu]=false
	c.keyDownMap[keyNextTrack]=false

	return c
}


func (c *controler) changeAngle(dir float64) {
	c.angle += c._angleLerp * dir

	if c.angle > 2*math.Pi || c.angle < -2*math.Pi{
		c.angle = math.Abs(c.angle)-2*math.Pi
	}
}

func (c *controler) changePower(dir float64) {
	if c.power < c._maxPower && dir > 0 {
		c.power += c._powerLerp
	}
	if c.power > defaultPowerLerp && dir < 0 {
		c.power -= c._powerLerp
	}
}

func hitKeyDown(hitKey ebiten.Key,b *ball) bool {
	if ebiten.IsKeyPressed(hitKey) {
		if b.controls.keyDownMap[hitKey] == false{
			b.controls.keyDownMap[hitKey] = true
			return true
		}
		return false
	}
	b.controls.keyDownMap[hitKey] = false
	return false
}

func mouseButtonDown(button ebiten.MouseButton) bool {
	if ebiten.IsMouseButtonPressed(button) {
		if mouseBtnIsDown == false {
			mouseBtnIsDown = true
			return true
		}
		return false
	}
	mouseBtnIsDown = false
	return false
}
