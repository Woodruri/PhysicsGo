package main

import (
	"math"
)

const DEFAULT_DRAG float64 = 0.1

func Vector2 struct {
	X, Y float64
}

type Body struct {
	Pos 		Vector2
	Vel 		Vector2
	Acc 		Vector2
	Mass 		float64
	Radius 		float64
	Restitution float64
}

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

func main() {
	fmt.Println(quote.Go())
}