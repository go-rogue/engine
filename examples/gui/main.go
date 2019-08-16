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

	var dt float32

	// Load the font tileset
	font := sprites.TileSetProperties{Filename: "arial10x10.png", Codec: sprites.LayoutTcod, W: 10, H: 10}

	// Initiate a Raylib windowed virtual Console for drawing to
	console := gui.NewRaylibConsole(screenWidth, screenHeight, 60, "example: gui", font, false)

	iGui := gui.NewGui(console)

	console.SetChar('a', cardinal.Position{1, 1}, rl.Red, rl.Color{})
	console.PutChar('b', cardinal.Position{2, 1})
	console.Print(cardinal.Position{1, 2}, "Hello World!")
	console.Print(cardinal.Position{1, 3}, "<%FG:colour_name>Text...<%/> Not Command Wrapped Text... <Help?>")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(console.GetDefaultBackground())

		dt = rl.GetFrameTime()

		gui.UpdateMouseStatus()
		iGui.UpdateWidgets(dt)
		console.Draw(dt)

		rl.EndDrawing()
	}

	// Clean up
	console.Unload()

	rl.CloseWindow()
}
