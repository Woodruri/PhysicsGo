package main

import (
	"image"
	"image/color"
	"os"
	"time"
)

const (
	// PHYSICS
	DEFAULT_DRAG float64 = 0.1
	GRAVITY float64 = 9.8
	BOUNCE_COEFFICIENT float64 = .8

	// SCREEN
	SCREEN_WIDTH  = 480
	SCREEN_HEIGHT = 320 
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

/******************************************************************

DRAWING FUNCS

******************************************************************/

func drawCircle(img *image.RGBA, circle *Circle) {
	for y := int(circle.Bod.Pos.Y) - int(circle.Radius) ; y <= int(circle.Bod.Pos.Y) ; y++ {
		for x := int(circle.Bod.Pos.X) - int(circle.Radius) ; x <= int(circle.Bod.Pos.X) ; x++ {
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
	canvas := image.NewRGBA(image.Rect(0, 0, SCREEN_HEIGHT, SCREEN_HEIGHT))

	// the BALL
	ball := &Circle{
		Bod: Body{
			Pos: Vector2{X: 20, Y: 20},
			Vel: Vector2{X: 5, Y: 5},
			Acc: Vector2{X: 0, Y: 0},
			Restitution: 10,
			Mass: 5,
		},
		Radius: 30,
		Clr: color.RGBA{100, 100, 255, 250},
	}
	
	// Time
	var dt float64 = 1.0 / 60.0
	
	//Main loop
	for {
		
		startTime := time.Now()
		
		//Physics Update
		// Gravity
		ball.Bod.Vel.Y += GRAVITY * dt
		ball.Bod.Vel.X += GRAVITY * dt

		// Floor Bounce
		if ball.Bod.Pos.Y >= float64(SCREEN_HEIGHT) - ball.Radius {
			ball.Bod.Pos.Y = float64(SCREEN_HEIGHT) - ball.Radius
			ball.Bod.Vel.Y = -ball.Bod.Vel.Y * BOUNCE_COEFFICIENT
		}

		//Rendering
		// Clear canvas
		clearCanvas(canvas)

		// Draw Ball
		drawCircle(canvas, ball)
		
		// Pipe to window
		os.Stdout.Write(canvas.Pix)

		// Wait 1/60 sec
		elapsedTime := time.Since(startTime)
		time.Sleep((time.Second / 60.0) - elapsedTime)
	}

}