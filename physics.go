package main

import (
	"image"
	"image/color"
	"os"
	"time"
	"math/rand"
)

const (
	// PHYSICS
	DEFAULT_DRAG float64 = 0.1
	GRAVITY float64 = 98
	BOUNCE_COEFFICIENT float64 = .95
	BOUNCE_DIR_COEFFICIENT float64 = 15

	// SCREEN
	SCREEN_WIDTH  = 1080
	SCREEN_HEIGHT = 720 

	// OTHER
	FRAMERATE = 60.0
	TOTAL_RUNS = 10000
)

/******************************************************************

STRUCTS

******************************************************************/


type Vector2 struct {
	X, Y float64
}

type Body struct {
	Pos 		Vector2
	Vel 		Vector2
	Acc 		Vector2
	Mass 		float64
	Restitution float64 // bounciness
}

type Circle struct {
	Bod         Body
	Radius 		float64
	Clr 		color.RGBA
}

/******************************************************************

PHYSICS FUNCS

******************************************************************/

func (body *Body) ApplyForce(force Vector2) {
	body.Acc.X += force.X / body.Mass
	body.Acc.Y += force.Y / body.Mass
}

func (body *Body) Update(dt float64) {
	body.Vel.X += body.Acc.X * dt
	body.Vel.Y += body.Acc.Y * dt

	body.Pos.X += body.Vel.X * dt
	body.Pos.Y += body.Vel.Y * dt

	body.Acc.X = 0
	body.Acc.Y = 0
}

// Handles collisions on borders
func (circ *Circle) Bounce(dt float64) {
	randomBool := rand.Intn(2) == 1

	if circ.Bod.Pos.Y >= float64(SCREEN_HEIGHT) - circ.Radius {
		circ.Bod.Pos.Y = float64(SCREEN_HEIGHT) - circ.Radius
		circ.Bod.Vel.Y = -circ.Bod.Vel.Y * BOUNCE_COEFFICIENT
		if randomBool {
			circ.Bod.Vel.X += rand.NormFloat64() * (BOUNCE_COEFFICIENT * -1)
		} else {
			circ.Bod.Vel.X += rand.NormFloat64() * BOUNCE_COEFFICIENT
		}
	}
	if circ.Bod.Pos.X <= 0 + circ.Radius || 
	circ.Bod.Pos.X >= float64(SCREEN_WIDTH) - circ.Radius {
		circ.Bod.Vel.X = -circ.Bod.Vel.X * BOUNCE_COEFFICIENT

		if circ.Bod.Pos.X >= float64(SCREEN_WIDTH) - circ.Radius {
			circ.Bod.Pos.X = float64(SCREEN_WIDTH) - circ.Radius
		} else if circ.Bod.Pos.X <= 0 + circ.Radius {
			circ.Bod.Pos.X = 0 + circ.Radius
		}
	}
}

/******************************************************************

DRAWING FUNCS

******************************************************************/

func drawCircle(img *image.RGBA, circle *Circle) {
	for y := int(circle.Bod.Pos.Y - circle.Radius) ; y <= int(circle.Bod.Pos.Y + circle.Radius) ; y++ {
		for x := int(circle.Bod.Pos.X - circle.Radius) ; x <= int(circle.Bod.Pos.X + circle.Radius) ; x++ {
			if x >= 0 && x < SCREEN_WIDTH && y >= 0 && y < SCREEN_HEIGHT {
				dx := x - int(circle.Bod.Pos.X)
				dy := y - int(circle.Bod.Pos.Y)
				if (dx*dx) + (dy*dy) <= int(circle.Radius * circle.Radius) {
					img.SetRGBA(int(x), int(y), circle.Clr)
				}
			}
		}
	}
}

func clearCanvas(img *image.RGBA) {
	for i := 0 ; i < len(img.Pix); i += 4 {
		img.Pix[i]   = 0   // R
		img.Pix[i+1] = 0   // G
		img.Pix[i+2] = 0   // B
		img.Pix[i+3] = 255 // A
	}
}

/******************************************************************

MAIN

******************************************************************/

func main() {
	//Setup
	// Canvas
	canvas := image.NewRGBA(image.Rect(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT))

	// the BALL
	ball := &Circle{
		Bod: Body{
			Pos: Vector2{X: 100, Y: 100},
			Vel: Vector2{X: 900, Y: 200},
			Acc: Vector2{X: 0, Y: 0},
			Restitution: 10,
			Mass: 5,
		},
		Radius: 30,
		Clr: color.RGBA{100, 100, 255, 250},
	}
	
	// Time
	var dt float64 = 1.0 / float64(FRAMERATE)
	targetFrameTime := time.Second / FRAMERATE
	
	//Main loop
	for {
		startTime := time.Now()
		
		//Physics Update
		// Gravity and Floor Bounce
		ball.Bod.ApplyForce(Vector2{X: 0, Y: GRAVITY * ball.Bod.Mass * 10})
		ball.Bod.Update(dt)
		ball.Bounce(dt)

		//Rendering
		// Clear canvas
		clearCanvas(canvas)

		// Draw Ball
		drawCircle(canvas, ball)
		
		// Pipe to window
		os.Stdout.Write(canvas.Pix)
		os.Stdout.Sync()

		// Wait framerate
		elapsedTime := time.Since(startTime)
		if elapsedTime < targetFrameTime {
			time.Sleep(targetFrameTime - elapsedTime)
		}
	}
}