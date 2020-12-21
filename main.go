package main

import (
	"fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	// DefaultScreenWidth ...
	DefaultScreenWidth int32 = 800
	// DefaultScreenHeight ...
	DefaultScreenHeight int32 = 500
	// DefaultFullScreenWidth ...
	DefaultFullScreenWidth int32 = 1280
	// DefaultFullScreenHeight ...
	DefaultFullScreenHeight int32 = 720
	// DefaultFont ...
	DefaultFont string = "resource/fontPixel.ttf"
)

type game struct {
	Name              string
	FullScreen        bool
	ResulutionScaling float32
	Font              rl.Font
	CurrentScene      string
	CursorColor       rl.Color
}

// Game ...
var Game = game{
	Name:              "The Cubes",
	FullScreen:        false,
	ResulutionScaling: 50.0,
	CurrentScene:      "menu",
	CursorColor:       rl.White,
}

func main() {
	// Set Logging
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprint(Game.Name, " >>> "))

	// Init Raylib
	log.Println("Starting and Setting Raylib...")
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(DefaultScreenWidth, DefaultScreenHeight, Game.Name)
	rl.SetTargetFPS(60)
	rl.HideCursor()
	Game.Font = rl.LoadFont(DefaultFont)
	Game.ResulutionScaling = float32((int32(rl.GetScreenWidth()) + int32(rl.GetScreenHeight())) / 25)
	Game.FullScreen = false

	// Game Loop
	log.Println("Starting Game Loop...")
	if Game.CurrentScene == "menu" {
		log.Println("Starting Scene: Menu...")
	}
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// F11 key fullscreen toggling
		if rl.IsKeyPressed(rl.KeyF11) {
			switch Game.FullScreen {
			case true:
				log.Println("Switching to windowed mode...")
				rl.CloseWindow()
				rl.SetConfigFlags(rl.FlagVsyncHint)
				rl.InitWindow(DefaultScreenWidth, DefaultScreenHeight, Game.Name)
				rl.SetTargetFPS(60)
				rl.HideCursor()
				Game.Font = rl.LoadFont(DefaultFont)
				Game.ResulutionScaling = float32((int32(rl.GetScreenWidth()) + int32(rl.GetScreenHeight())) / 25)
				Game.FullScreen = false

			case false:
				log.Println("Switching to fullscreen mode...")
				rl.CloseWindow()
				rl.SetConfigFlags(rl.FlagVsyncHint)
				rl.SetConfigFlags(rl.FlagWindowUndecorated)
				rl.InitWindow(DefaultFullScreenWidth, DefaultFullScreenHeight, Game.Name)
				rl.SetTargetFPS(60)
				rl.HideCursor()
				Game.Font = rl.LoadFont(DefaultFont)
				Game.ResulutionScaling = float32((int32(rl.GetScreenWidth()) + int32(rl.GetScreenHeight())) / 25)
				Game.FullScreen = true
			}
		}

		// Scenes
		switch Game.CurrentScene {
		case "menu":
			SceneMenu()
		}

		rl.EndDrawing()
	}

	// The End.
	rl.CloseWindow()
}
