package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-rogue/engine/gui"
	"github.com/go-rogue/engine/scenes"
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

	director := scenes.NewDirector(NewButtonScene(iGui))

	updateWindowTitle := false

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyO) {
			director.ChangeState(NewSkeletonScene(iGui))
			updateWindowTitle = true

		} else if rl.IsKeyPressed(rl.KeyP) {
			director.ChangeState(NewButtonScene(iGui))
			updateWindowTitle = true
		}

		if updateWindowTitle {
			rl.SetWindowTitle(fmt.Sprintf("example: %s", director.PeekState().GetName()))
			updateWindowTitle = false
		}

		rl.BeginDrawing()

		dt = rl.GetFrameTime()
		gui.UpdateMouseStatus()
		director.PeekState().Tick(dt)

		iGui.UpdateWidgets(dt)
		iGui.RenderWidgets()
		console.Draw(dt)

		rl.EndDrawing()
	}

	// Clean up
	console.Unload()

	rl.CloseWindow()
}
