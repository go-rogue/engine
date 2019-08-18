package main

import (
	"fmt"
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

	console.PutCharEx('a', cardinal.Position{1, 1}, rl.Red, rl.Color{})
	console.SetChar('b', cardinal.Position{2, 1})
	console.Print(cardinal.Position{1, 2}, "Hello World!")
	console.Print(cardinal.Position{1, 3}, "<%FG:red>Text...<%/> Not Command Wrapped Text... <Help?>")
	console.Print(cardinal.Position{1, 4}, "<%FG:white,BG:red>Text, <%FG:blue>this has depth<%/> this doesn't. <%BG:blue>blue bg<%/>...<%/> back to normal")

	iGui.NewButton(cardinal.Position{1, 10}, 11, 3, "hello", "", gui.SingleWallBorder, func(w *gui.Button, userData interface{}) {
		if w.GetLabel() == "hello" {
			w.SetLabel("goodbye")
		} else {
			w.SetLabel("hello")
		}
	}, nil)

	iGui.NewButton(cardinal.Position{13, 10}, 11, 3, "hello", "", gui.SingleWallBorder, func(w *gui.Button, userData interface{}) {
		if w.GetLabel() == "hello" {
			w.SetLabel("goodbye")
		} else {
			w.SetLabel("hello")
		}
	}, nil)

	iGui.NewButton(cardinal.Position{23, 15}, 21, 7, "Click to Disable", "", gui.SingleWallBorder, func(w *gui.Button, userData interface{}) {
		w.SetDisabled(true)
		w.SetUserData(w.GetUserData().(int) + 1)
		w.SetLabel(fmt.Sprintf("Clicked %d", w.GetUserData().(int)))
	}, 0)

	iGui.NewButton(cardinal.Position{1, 15}, 21, 7, "Click Me", "", gui.SingleWallBorder, func(w *gui.Button, userData interface{}) {
		w.SetUserData(w.GetUserData().(int) + 1)
		w.SetLabel(fmt.Sprintf("Clicked %d", w.GetUserData().(int)))
	}, 0)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(console.GetDefaultBackground())

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
