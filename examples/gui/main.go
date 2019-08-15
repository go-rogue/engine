package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-rogue/engine/cardinal"
	"github.com/go-rogue/engine/gui"
	"github.com/go-rogue/engine/sprites"
)

const (
	screenWidth  = 800
	screenHeight = 450
)

func main() {

	// Load the font tileset
	font := sprites.TileSetProperties{Filename: "arial10x10.png", Codec: sprites.LayoutTcod, W: 10, H: 10}

	// Initiate a Raylib windowed virtual Console for drawing to
	console := gui.NewRaylibConsole(screenWidth, screenHeight, 60, "example: gui", font, false)

	// gui.NewGui(console)

	console.SetChar('a', cardinal.Position{1, 1}, rl.Red, rl.Color{})
	console.SetChar('b', cardinal.Position{2, 1}, rl.Red, rl.Color{})
	console.SetChar('c', cardinal.Position{3, 1}, rl.Red, rl.Color{})
	console.SetChar('@', cardinal.Position{4, 1}, rl.Red, rl.Color{})

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		console.Draw()

		rl.EndDrawing()
	}

	// Clean up
	console.Unload()

	rl.CloseWindow()
}
