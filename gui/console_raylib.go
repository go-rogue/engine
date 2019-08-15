package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-rogue/engine/sprites"
)

type RaylibConsole struct {
	Console
	tileset *sprites.Tileset
}

//
// Create a new Virtual console wrapped in a Raylib window; only one of these should be used per application.
//
// For a Windowed console the width/height of Console is equivalent to columns/rows.
//
// @todo check and return error if division by zero when tile width/height is zero
//
func NewRaylibConsole(w, h uint, fps uint, title string, fontProps sprites.TileSetProperties, fullscreen bool) *RaylibConsole {

	rl.InitWindow(int32(w), int32(h), title)
	rl.SetTargetFPS(int32(fps))

	if fullscreen == true {
		rl.ToggleFullscreen()
	}

	ts := sprites.NewTileSetFromProps(fontProps)

	ret := &RaylibConsole{
		Console: Console{width: w / ts.GetTileWidth(), height: h / ts.GetTileHeight()},
		tileset: ts,
	}

	ret.init() // this is found on the parent Console

	return ret
}

func (c RaylibConsole) Draw() {
	for pos, cell := range *c.GetData() {
		c.tileset.GetSpriteForChar(cell.char).Draw(pos, cell.fg, cell.bg)
	}
}

func (c RaylibConsole) Unload() {
	c.tileset.Unload()
}
