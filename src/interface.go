package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(960, 640, "Minecraft Edition Pixel War")
	defer rl.CloseWindow()
	rl.SetTargetFPS(120)

	// Charger l'image PNG en tant qu'arrière-plan
	background := rl.LoadTexture("background.jpg")
	logo := rl.LoadTexture("logo.png")

	// Calculer les coordonnées x et y pour centrer le logo
	screenWidth := int32(rl.GetScreenWidth())
	screenHeight := int32(rl.GetScreenHeight())
	logoX := (screenWidth - int32(logo.Width)) / 2
	logoY := (screenHeight - int32(logo.Height)) / 2

	// Définir les propriétés du bouton "Play"
	playButtonWidth := int32(200)
	playButtonHeight := int32(50)
	playButtonX := (screenWidth - playButtonWidth) / 2
	playButtonY := logoY + int32(logo.Height) + 20
	playButtonColor := rl.NewColor(150, 150, 150, 255) // Couleur du bouton "Play"

	// Définir les propriétés du bouton "Quitter"
	quitButtonWidth := int32(200)
	quitButtonHeight := int32(50)
	quitButtonX := (screenWidth - quitButtonWidth) / 2
	quitButtonY := playButtonY + playButtonHeight + 20
	quitButtonColor := rl.NewColor(150, 150, 150, 255) // Couleur du bouton "Quitter"

	for !rl.WindowShouldClose() {
		// Détecter le clic de la souris
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			mousePos := rl.GetMousePosition()

			// Vérifier si le clic est sur le bouton "Quitter"
			if mousePos.X >= float32(quitButtonX) && mousePos.X <= float32(quitButtonX+quitButtonWidth) &&
				mousePos.Y >= float32(quitButtonY) && mousePos.Y <= float32(quitButtonY+quitButtonHeight) {
				// Fermer la fenêtre si le bouton "Quitter" est cliqué
				rl.CloseWindow()
			}
		}

		rl.BeginDrawing()

		// Importer le background / logo
		rl.DrawTexture(background, 0, 0, rl.RayWhite)
		rl.DrawTexture(logo, logoX, logoY, rl.RayWhite)

		// Dessiner le bouton "Play"
		rl.DrawRectangle(playButtonX, playButtonY, playButtonWidth, playButtonHeight, playButtonColor)
		rl.DrawText("Jouer", int32(playButtonX)+10, playButtonY+10, 30, rl.Black)

		// Dessiner le bouton "Quitter"
		rl.DrawRectangle(quitButtonX, quitButtonY, quitButtonWidth, quitButtonHeight, quitButtonColor)
		rl.DrawText("Quitter", int32(quitButtonX)+10, quitButtonY+10, 30, rl.Black)

		rl.EndDrawing()
	}

	// Libérer la texture de l'arrière-plan
	rl.UnloadTexture(background)
	rl.UnloadTexture(logo)
}
