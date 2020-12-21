package main

import (
	// "net/url"

	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// SceneMenu ...
func SceneMenu() {
	if CurrentScene != "menu" {
		log.Println("Starting Scene: Menu...")
		CurrentScene = "menu"
	}

	rl.BeginDrawing()

	// var url url.URL // := url.URL{Scheme: "ws", Host: addr, Path: "/"}

	// Background
	rl.ClearBackground(rl.Gray)

	// UI
	rl.DrawTextEx(Font, GameName, rl.NewVector2(float32(rl.GetScreenWidth()/2), float32(rl.GetScreenHeight()/2)), float32(ResulutionScaling), 0, rl.White)

	rl.EndDrawing()
}
