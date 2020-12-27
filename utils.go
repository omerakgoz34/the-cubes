package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// DrawCustomCursor = draws custom cursor to the screen instead of default OS cursor
func DrawCustomCursor() {
	cursorScale := int32(Game.ResulutionScaling / 10)
	rl.DrawRectangle(rl.GetMouseX(), rl.GetMouseY(), cursorScale, cursorScale, Game.CursorColor)
	rl.DrawRectangle(rl.GetMouseX()+cursorScale, rl.GetMouseY()+cursorScale, cursorScale, cursorScale, Game.CursorColor)
	rl.DrawRectangle(rl.GetMouseX()+cursorScale*2, rl.GetMouseY()+cursorScale*2, cursorScale, cursorScale, Game.CursorColor)
	rl.DrawRectangle(rl.GetMouseX(), rl.GetMouseY()+cursorScale, cursorScale, cursorScale, Game.CursorColor)
	rl.DrawRectangle(rl.GetMouseX(), rl.GetMouseY()+cursorScale*2, cursorScale, cursorScale, Game.CursorColor)
	rl.DrawRectangle(rl.GetMouseX(), rl.GetMouseY()+cursorScale*3, cursorScale, cursorScale, Game.CursorColor)
	rl.DrawRectangle(rl.GetMouseX()+cursorScale, rl.GetMouseY()+cursorScale*2, cursorScale, cursorScale, Game.CursorColor)
}
