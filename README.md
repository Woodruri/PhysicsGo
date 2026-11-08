# PhysicsGo

A small 3D physics sandbox written in Go with [raylib-go](https://github.com/gen2brain/raylib-go). A ball falls under gravity and bounces around inside a 10×10×10 wireframe box. You can push it and make it jump with the keyboard.

## Requirements

- Go 1.24 or newer
- raylib-go's platform requirements. On Windows, raylib-go can run without cgo, but then it needs `raylib.dll` next to the executable or on your `PATH`. See the [raylib-go README](https://github.com/gen2brain/raylib-go#requirements) for the details on each OS.

## Running

```sh
go mod tidy      
go run physics.go
```

The program opens a 1080×720 window and runs at 60 FPS.

## Controls

| Key           | Action                            |
|---------------|-----------------------------------|
| Arrow keys    | Push the ball along the X / Z axes |
| Space         | Jump (adds upward velocity)       |
| Mouse wheel   | Zoom (the camera orbits on its own) |
| Esc           | Quit                              |

## How it works

Everything lives in [physics.go](physics.go):

- **`Ball.Update(dt)`** applies gravity to the Y velocity, moves the ball by `velocity × dt`, and then checks the walls. When the ball goes past a wall on any axis, it is clamped back inside and that velocity component is reversed and scaled down by `BOUNCE`.
- **`main`** sets up the window and an orbiting camera, then runs a loop each frame: input, physics step, render.

Y points up. Positions and velocities are in meters and m/s.

### Tunable constants

| Constant  | Default | Meaning                                         |
|-----------|---------|-------------------------------------------------|
| `GRAVITY` | `9.8`   | Downward acceleration (m/s²)                    |
| `BOUNCE`  | `0.9`   | Fraction of speed kept after hitting a wall     |
| `HALF`    | `5.0`   | Half the width of the box (the box is `2*HALF` per side) |
