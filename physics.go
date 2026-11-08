package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	// PHYSICS
	GRAVITY = 9.8
	BOUNCE  = 0.9

	// WORLD
	HALF    = 5.0 // half the size of the box
)

/******************************************************************
STRUCTS
******************************************************************/

type Ball struct {
	Pos, Vel rl.Vector3
	Radius   float32
}

/******************************************************************
PHYSICS FUNCS
******************************************************************/

func (b *Ball) Update(dt float32) {
	b.Vel.Y -= GRAVITY * dt
	b.Pos = rl.Vector3Add(b.Pos, rl.Vector3Scale(b.Vel, dt))

	// Handles collisions on walls
	bounce := func(p, v *float32) {
		lim := HALF - b.Radius
		if *p > lim  { *p = lim;  *v = -*v * BOUNCE }
		if *p < -lim { *p = -lim; *v = -*v * BOUNCE }
	}
	bounce(&b.Pos.X, &b.Vel.X)
	bounce(&b.Pos.Y, &b.Vel.Y)
	bounce(&b.Pos.Z, &b.Vel.Z)
}

/******************************************************************
MAIN
******************************************************************/

func main() {
	//Setup
	// Window
	rl.InitWindow(1080, 720, "PhysicsGo 3D")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// Camera
	cam := rl.Camera3D{
		Position:   rl.NewVector3(12, 8, 12),
		Target:     rl.NewVector3(0, 0, 0),
		Up:         rl.NewVector3(0, 1, 0),
		Fovy:       45,
		Projection: rl.CameraPerspective,
	}

	// the BALL
	ball := &Ball{Pos: rl.NewVector3(0, 3, 0), Vel: rl.NewVector3(6, 0, 3), Radius: 0.5}

	//Main loop
	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		//Input
		// Camera
		rl.UpdateCamera(&cam, rl.CameraOrbital)

		// Jump
		if rl.IsKeyPressed(rl.KeySpace) { ball.Vel.Y += 8 }

		// Push ball
		if rl.IsKeyDown(rl.KeyLeft)     { ball.Vel.X -= 10 * dt }
		if rl.IsKeyDown(rl.KeyRight)    { ball.Vel.X += 10 * dt }
		if rl.IsKeyDown(rl.KeyUp)       { ball.Vel.Z -= 10 * dt }
		if rl.IsKeyDown(rl.KeyDown)     { ball.Vel.Z += 10 * dt }

		//Physics Update
		ball.Update(dt)

		//Rendering
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.BeginMode3D(cam)

		// Draw Box
		rl.DrawCubeWires(rl.NewVector3(0, 0, 0), 2*HALF, 2*HALF, 2*HALF, rl.Gray)

		// Draw Ball
		rl.DrawSphere(ball.Pos, ball.Radius, rl.NewColor(100, 100, 255, 255))

		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
