package main

import (
	"fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// GameName ...
const GameName string = "The Cubes"

// ScreenWidth ...
const ScreenWidth int32 = 800

// ScreenHeight ...
const ScreenHeight int32 = 500

// FullScreenWidth ...
const FullScreenWidth int32 = 1280

// FullScreenHeight ...
const FullScreenHeight int32 = 720

// FullScreen ...
var FullScreen bool = false

// Font ...
var Font rl.Font

// CurrentScene ...
var CurrentScene string = "menu"

func main() {
	// Set Logging
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprint(GameName, " >>> "))

	// Init Raylib
	log.Println("Starting Raylib...")
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(ScreenWidth, ScreenHeight, GameName)
	rl.SetTargetFPS(60)

	// Load Font
	Font = rl.LoadFont("bscePixel.ttf")

	// Game Loop
	log.Println("Starting Game Loop...")
	if CurrentScene == "menu" {
		log.Println("Starting Scene: Menu...")
	}
	for !rl.WindowShouldClose() {

		// F11 key fullscreen toggling
		if rl.IsKeyPressed(rl.KeyF11) {
			switch FullScreen {
			case true:
				log.Println("Switching to windowed mode...")
				rl.CloseWindow()
				rl.SetConfigFlags(rl.FlagVsyncHint)
				rl.InitWindow(ScreenWidth, ScreenHeight, GameName)
				rl.SetTargetFPS(60)
				FullScreen = false

			case false:
				log.Println("Switching to fullscreen mode...")
				rl.CloseWindow()
				rl.SetConfigFlags(rl.FlagVsyncHint)
				rl.SetConfigFlags(rl.FlagWindowUndecorated)
				rl.InitWindow(FullScreenWidth, FullScreenHeight, GameName)
				rl.SetTargetFPS(60)
				FullScreen = true
			}
		}

		// Scenes
		switch CurrentScene {
		case "menu":
			SceneMenu()
		}
	}

	// The End.
	rl.CloseWindow()
}
