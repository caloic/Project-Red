package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"time"
)

func main() {
	rl.InitWindow(960, 640, "Je suis une grosse salope !")
	defer rl.CloseWindow()
	rl.SetTargetFPS(120)

	// Charger l'image PNG en tant qu'arrière-plan
	background := rl.LoadTexture("background.jpg")

	// Couleurs personnalisées
	textColor := rl.RayWhite

	// Initialisation du générateur de nombres aléatoires avec une graine unique
	rand.Seed(time.Now().UnixNano())

	// Temps en millisecondes pour changer de couleur (ultra rapide)
	colorChangeInterval := time.Millisecond * 100
	lastColorChangeTime := time.Now()

	// Texte
	message := "Vous devez me sucer pour gagner la partie"
	fontSize := int32(20)

	// Rectangle pour le bouton "Play"
	playButtonRect := rl.NewRectangle(400, 350, 160, 40) // Ajustez les valeurs selon votre mise en page

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// Dessiner l'arrière-plan avec l'image PNG
		rl.DrawTexture(background, 0, 0, rl.RayWhite)

		// Vérifier si le temps écoulé est supérieur à l'intervalle de changement de couleur
		if time.Since(lastColorChangeTime) > colorChangeInterval {
			// Génération de couleurs aléatoires pour le texte
			textColor.R = uint8(rand.Intn(256))
			textColor.G = uint8(rand.Intn(256))
			textColor.B = uint8(rand.Intn(256))

			// Mettre à jour le temps du dernier changement de couleur
			lastColorChangeTime = time.Now()
		}

		// Mesurer la largeur du texte pour le centrer
		textWidth := rl.MeasureText(message, fontSize)
		x := (int32(rl.GetScreenWidth()) - textWidth) / 2
		y := (int32(rl.GetScreenHeight()) - fontSize) / 2

		// Afficher le texte centré
		rl.DrawText(message, x, y, fontSize, textColor)

		// Dessiner le bouton "Play"
		rl.DrawRectangleLinesEx(playButtonRect, 2, rl.RayWhite)
		rl.DrawText("Play", int32(playButtonRect.X)+10, int32(playButtonRect.Y)+7, fontSize-2, rl.RayWhite)

		// Vérifier si la souris a cliqué sur le bouton "Play"
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), playButtonRect) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			// Ouvrir le navigateur Web avec le lien YouTube
			rl.OpenURL("https://www.youtube.com/") // Remplacez par le lien de votre choix
		}

		rl.EndDrawing()
	}

	// Libérer la texture de l'arrière-plan
	rl.UnloadTexture(background)
}
