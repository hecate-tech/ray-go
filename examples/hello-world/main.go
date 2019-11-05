package main

import "github.com/hecate-tech/ray-go/ray"

func main() {
	ray.InitWindow(800, 450, "Hello Raylib")
	defer ray.CloseWindow()

	ray.SetTargetFPS(60)

	for !ray.WindowShouldClose() {
		ray.BeginDrawing()

		ray.ClearBackground(ray.White)

		ray.DrawText("Hello, World!", 50, 50, 50, ray.Black)

		ray.EndDrawing()
	}
}
