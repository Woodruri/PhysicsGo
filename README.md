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
