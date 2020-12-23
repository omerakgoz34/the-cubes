package main

import (
	"fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kbinani/screenshot"
)

const (
	// DefaultFont ...
	DefaultFont string = "res/fontPixel.ttf"
)

type game struct {
	Name              string
	FullScreen        bool
	ScreenWidth       int32
	ScreenHeight      int32
	FullScreenWidth   int32
	FullScreenHeight  int32
	ResulutionScaling float32
	Font              rl.Font
	CurrentScene      string
	CursorColor       rl.Color
}

// Game ...
var Game = game{
	Name:              "The Cubes",
	FullScreen:        false,
	ScreenWidth:       960,
	ScreenHeight:      540,
	FullScreenWidth:   960,
	FullScreenHeight:  540,
	ResulutionScaling: 50.0,
	CurrentScene:      "menu",
	CursorColor:       rl.White,
}

func main() {
	// Set Logging
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprint(Game.Name, " >>> "))

	// Get Screen Resolution
	Game.FullScreenWidth = int32(screenshot.GetDisplayBounds(0).Dx())
	Game.FullScreenHeight = int32(screenshot.GetDisplayBounds(0).Dy())

	// Init Raylib
	log.Println("Starting and Setting Raylib...")
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(Game.ScreenWidth, Game.ScreenHeight, Game.Name)
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
				log.Println("Switching to Windowed Mode...")
				rl.CloseWindow()
				rl.SetConfigFlags(rl.FlagVsyncHint)
				rl.InitWindow(Game.ScreenWidth, Game.ScreenHeight, Game.Name)
				rl.SetTargetFPS(60)
				rl.HideCursor()
				Game.Font = rl.LoadFont(DefaultFont)
				Game.ResulutionScaling = float32((int32(rl.GetScreenWidth()) + int32(rl.GetScreenHeight())) / 25)
				Game.FullScreen = false

			case false:
				log.Println("Switching to Fullscreen Mode...")
				rl.CloseWindow()
				rl.SetConfigFlags(rl.FlagVsyncHint)
				rl.SetConfigFlags(rl.FlagWindowUndecorated)
				rl.InitWindow(Game.FullScreenWidth, Game.FullScreenHeight, Game.Name)
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
