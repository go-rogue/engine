package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-rogue/engine/gui"
	"github.com/go-rogue/engine/sprites"
)

const (
	screenWidth  = 800
	screenHeight = 450
)

func main() {
	var dt float32

	// Load the font tileset
	font := sprites.TileSetProperties{Filename: "arial10x10.png", Codec: sprites.LayoutTcod, W: 10, H: 10}

	// Initiate a Raylib windowed virtual Console for drawing to
	console := gui.NewRaylibConsole(screenWidth, screenHeight, 60, "example: skeleton", font, false)

	iGui := gui.NewGui(console)

	// ...

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		dt = rl.GetFrameTime()
		gui.UpdateMouseStatus()
		iGui.UpdateWidgets(dt)
		iGui.RenderWidgets()
		console.Draw(dt)

		rl.EndDrawing()
	}

	// Clean up
	console.Unload()

	rl.CloseWindow()
}
