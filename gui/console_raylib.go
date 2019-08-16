package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-rogue/engine/cardinal"
	"github.com/go-rogue/engine/sprites"
	"math"
)

type RaylibConsole struct {
	Console
	tileset *sprites.Tileset
}

func PositionFromVec2(vector2 rl.Vector2) cardinal.Position {
	return cardinal.Position{X: int(vector2.X), Y: int(vector2.Y)}
}

//
// Create a new Virtual console wrapped in a Raylib window; only one of these should
// be used per application. In libtcod parlance this is a root console.
//
// For a Windowed console the width/height of Console is equivalent to columns/rows.
//
// @todo check and return error if division by zero when tile width/height is zero
//
func NewRaylibConsole(w, h uint, fps uint, title string, fontProps sprites.TileSetProperties, fullscreen bool) *RaylibConsole {
	// NOTE: Textures and Sounds MUST be loaded after Window/Audio initialization
	rl.InitWindow(int32(w), int32(h), title)
	rl.SetTargetFPS(int32(fps))

	if fullscreen == true {
		rl.ToggleFullscreen()
	}

	ts := sprites.NewTileSetFromProps(fontProps)

	ret := &RaylibConsole{
		Console: *NewVirtualConsole(w/ts.GetTileWidth(), h/ts.GetTileHeight()),
		tileset: ts,
	}

	// Raylib has support for mouse input
	MouseStatus.Supported = true

	// Update the Mouse struct (this is because we wont just be handling raylib)
	UpdateMouseStatus = func() {
		mousePos := rl.GetMousePosition()
		position := PositionFromVec2(mousePos)
		position.X = int(math.Floor(float64(position.X / int(ts.GetTileWidth()))))
		position.Y = int(math.Floor(float64(position.Y / int(ts.GetTileHeight()))))

		MouseStatus.Pos = position
		MouseStatus.LButton = rl.IsMouseButtonDown(rl.MouseLeftButton)
		MouseStatus.MButton = rl.IsMouseButtonDown(rl.MouseMiddleButton)
		MouseStatus.RButton = rl.IsMouseButtonDown(rl.MouseRightButton)
		MouseStatus.LButtonPressed = rl.IsMouseButtonPressed(rl.MouseLeftButton)
		MouseStatus.MButtonPressed = rl.IsMouseButtonPressed(rl.MouseMiddleButton)
		MouseStatus.RButtonPressed = rl.IsMouseButtonPressed(rl.MouseRightButton)
		MouseStatus.WindowFocus = mousePos.X > 0 && mousePos.Y > 0 && mousePos.X < float32(rl.GetScreenWidth()) && mousePos.Y < float32(rl.GetScreenHeight())
		MouseStatus.Visible = !rl.IsCursorHidden()
	}

	ret.init() // this is found on the parent Console struct

	return ret
}

func (c RaylibConsole) Draw(dt float32) {
	for pos, cell := range *c.GetData() {
		c.tileset.GetSpriteForChar(cell.char).Draw(pos, cell.fg, cell.bg)
	}
}

func (c RaylibConsole) Unload() {
	c.tileset.Unload()
}
