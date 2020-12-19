package main

import (
	// "net/url"

	"log"

	rg "github.com/gen2brain/raylib-go/raygui"
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

	// Begin UI
	rg.Button(rl.NewRectangle(float32(rl.GetScreenWidth()/2-50), float32(rl.GetScreenHeight()/2-25), 100, 50), "Start")
	// End UI

	rl.EndDrawing()
}
