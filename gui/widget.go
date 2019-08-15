package gui

import rl "github.com/gen2brain/raylib-go/raylib"

type IWidget interface {
	SetGui(*Gui)
	GetGui() *Gui
}

type Widget struct {
	x, y, w, h int
	userData   interface{}
	tip        string
	mouseIn    bool
	mouseL     bool
	visible    bool
	back       rl.Color
	fore       rl.Color
	backFocus  rl.Color
	foreFocus  rl.Color
	gui        *Gui
}

type WidgetCallback func(w IWidget, userData interface{})
