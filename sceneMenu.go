package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	backgroundColor rl.Color = rl.SkyBlue
)

// SceneMenuInit ...
func SceneMenuInit() {

}

// SceneMenu ...
func SceneMenu() {
	if Game.CurrentScene != "menu" {
		log.Println("Starting Scene: Menu...")
		Game.CurrentScene = "menu"
	}

	// Background
	rl.ClearBackground(backgroundColor)

	// Title
	titleSize := rl.MeasureTextEx(Game.Font, Game.Name, float32(Game.ResulutionScaling*1.5), 0).X
	rl.DrawTextEx(Game.Font, Game.Name, rl.NewVector2(float32(rl.GetScreenWidth()/2)-titleSize/2, float32(rl.GetScreenHeight()/6)), Game.ResulutionScaling*1.5, 0, rl.White)

	// Basic Start Button
	startTitleSize := [2]float32{rl.MeasureTextEx(Game.Font, "Start", float32(Game.ResulutionScaling*1), 0).X, rl.MeasureTextEx(Game.Font, "Start", float32(Game.ResulutionScaling*1), 0).Y}
	rl.DrawRectangle(int32(float32((rl.GetScreenWidth()/2))-((startTitleSize[0]+10)/2)), int32(float32((rl.GetScreenHeight()/2))-((startTitleSize[1]+10)/2)), int32(startTitleSize[0]+10), int32(startTitleSize[1]+10), rl.Gray)
	rl.DrawTextEx(Game.Font, "Start", rl.NewVector2(float32(rl.GetScreenWidth()/2)-startTitleSize[0]/2, float32(rl.GetScreenHeight()/2)-startTitleSize[1]/2), Game.ResulutionScaling*1, 0, rl.White)

	// Floor
	rl.DrawRectangle(0, int32(rl.GetScreenHeight()-(rl.GetScreenHeight()/6)), int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()/10), rl.Green)
	rl.DrawRectangle(0, int32(rl.GetScreenHeight()-(rl.GetScreenHeight()/8)), int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Brown)

	// Custom Cursor
	DrawCustomCursor()
}
