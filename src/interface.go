package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(960, 640, "Minecraft Edition Pixel War")
	defer rl.CloseWindow()
	rl.SetTargetFPS(120)

	// Charger l'image PNG en tant qu'arrière-plan
	background := rl.LoadTexture("background.png")
	logo := rl.LoadTexture("logo.png")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// Impoter le backgroud / logo
		rl.DrawTexture(background, 0, 0, rl.RayWhite)
		rl.DrawTexture(logo, 100, 100, rl.RayWhite)
		rl.EndDrawing()
	}

	// Libérer la texture de l'arrière-plan
	rl.UnloadTexture(background)
	rl.UnloadTexture(logo)
}
