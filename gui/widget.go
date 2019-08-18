package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-rogue/engine/geom"
	"github.com/go-rogue/engine/sprites"
)

type BorderStyle struct {
	V, H, NE, SE, SW, NW uint
}

var SingleWallBorder = BorderStyle{
	sprites.TCOD_CHAR_VLINE, sprites.TCOD_CHAR_HLINE, sprites.TCOD_CHAR_NE, sprites.TCOD_CHAR_SE, sprites.TCOD_CHAR_SW, sprites.TCOD_CHAR_NW,
}

var ZeroWallBorder = BorderStyle{
	0, 0, 0, 0, 0, 0,
}

type IWidget interface {
	SetGui(*Gui)
	GetGui() *Gui
	Delete()
	GetWidth() uint
	SetWidth(width uint)
	GetHeight() uint
	SetHeight(height uint)
	GetPosition() geom.Position
	Move(pos geom.Position)
	IsVisible() bool
	SetVisible(visible bool)
	GetUserData() interface{}
	SetUserData(data interface{})
	SetDisabled(d bool)
	GetDisabled() bool
	GetTip() string
	SetTip(tip string)
	GetMouseIn() bool
	SetMouseIn(mouseIn bool)
	GetMouseL() bool
	SetMouseL(mouseL bool)
	ComputeSize()
	Update(w IWidget)
	Render(w IWidget)
	SetDefaultBackground(col, colFocus rl.Color)
	SetDefaultForeground(col, colFocus rl.Color)
	GetDefaultBackground() (col, colFocus rl.Color)
	GetDefaultForeground() (col, colFocus rl.Color)
	GetCurrentColors() (fore, back rl.Color)
	onMouseIn()
	onMouseOut()
	onButtonPress()
	onButtonRelease()
	onButtonClick()
	expand(width, height uint)
}

type Widget struct {
	pos          geom.Position
	w, h         uint
	userData     interface{}
	tip          string
	disabled     bool
	mouseIn      bool
	mouseL       bool
	visible      bool
	back         rl.Color
	fore         rl.Color
	backFocus    rl.Color
	foreFocus    rl.Color
	backDisabled rl.Color
	foreDisabled rl.Color
	gui          *Gui
}

type WidgetCallback func(w IWidget, userData interface{})

func (w *Widget) SetGui(g *Gui) {
	w.gui = g
}

func (w Widget) GetGui() *Gui {
	return w.gui
}

func (w *Widget) Delete() {
	w.gui.Unregister(w)
}

func (w Widget) GetWidth() uint {
	return w.w
}

func (w *Widget) SetWidth(width uint) {
	w.w = width
}

func (w Widget) GetHeight() uint {
	return w.h
}

func (w *Widget) SetHeight(height uint) {
	w.h = height
}

func (w Widget) GetPosition() geom.Position {
	return w.pos
}

func (w *Widget) Move(pos geom.Position) {
	w.pos = pos
}

func (w Widget) IsVisible() bool {
	return w.visible
}

func (w *Widget) SetVisible(visible bool) {
	w.visible = visible
}

func (w *Widget) GetUserData() interface{} {
	return w.userData
}

func (w *Widget) SetUserData(data interface{}) {
	w.userData = data
}

func (w *Widget) SetDisabled(d bool) {
	w.disabled = d
}

func (w Widget) GetDisabled() bool {
	return w.disabled
}

func (w Widget) GetTip() string {
	return w.tip
}

func (w *Widget) SetTip(tip string) {
	w.tip = tip
}

func (w Widget) GetMouseIn() bool {
	return w.mouseIn
}

func (w *Widget) SetMouseIn(mouseIn bool) {
	w.mouseIn = mouseIn
}

func (w Widget) GetMouseL() bool {
	return w.mouseL
}

func (w *Widget) SetMouseL(mouseL bool) {
	w.mouseL = mouseL
}

func (w *Widget) SetDefaultBackground(col, colFocus rl.Color) {
	w.back = col
	w.backFocus = colFocus
}

func (w *Widget) SetDefaultForeground(col, colFocus rl.Color) {
	w.fore = col
	w.foreFocus = colFocus
}

func (w *Widget) GetDefaultBackground() (col, colFocus rl.Color) {
	return w.back, w.backFocus
}

func (w *Widget) GetDefaultForeground() (col, colFocus rl.Color) {
	return w.fore, w.foreFocus
}

func (w *Widget) GetCurrentColors() (fore, back rl.Color) {
	if w.disabled {
		fore = w.foreDisabled
		back = w.backDisabled
	} else if w.mouseIn {
		fore = w.foreFocus
		back = w.backFocus
	} else {
		fore = w.fore
		back = w.back
	}

	return fore, back
}

//
// This gets called by the Gui and updates each widgets mouse meta
// w is the base Widget struct while iW is any of the inheriting
// widgets e.g Button.
//
func (w *Widget) Update(iW IWidget) {

	// If the console in use does not support the mouse then no point in updating
	// from it.
	if !w.gui.mouse.Supported {
		return
	}

	// Update Mouse In/Out/Focus if the cursor is visible
	if w.gui.mouse.Visible {
		iWPos := iW.GetPosition()
		if w.gui.mouse.Pos.X >= iWPos.X && w.gui.mouse.Pos.X < iWPos.X+int(iW.GetWidth()) &&
			w.gui.mouse.Pos.Y >= iWPos.Y && w.gui.mouse.Pos.Y < iWPos.Y+int(iW.GetHeight()) {
			if !iW.GetMouseIn() {
				iW.SetMouseIn(true)
				iW.onMouseIn()
			}
			if w.gui.focus != iW {
				w.gui.focus = iW
			}
		} else {
			if iW.GetMouseIn() {
				iW.SetMouseIn(false)
				iW.onMouseOut()
			}
			iW.SetMouseL(false)
			if iW == w.gui.focus {
				w.gui.focus = nil
			}
		}
	}

	// Update Mouse click/press/etc
	if iW.GetMouseIn() || (!w.gui.mouse.Visible && iW == w.gui.focus) {
		if w.gui.mouse.LButton && !iW.GetMouseL() {
			iW.SetMouseL(true)
			iW.onButtonPress()
		} else if !w.gui.mouse.LButton && iW.GetMouseL() {
			iW.onButtonRelease()
			w.gui.keyboardFocus = nil
			if iW.GetMouseL() {
				iW.onButtonClick()
			}
			iW.SetMouseL(false)
		} else if w.gui.mouse.LButtonPressed {
			w.gui.keyboardFocus = nil
			iW.onButtonClick()
		}
	}

}

//
// The following abstract methods will be filled in by each Widget
//

func (w *Widget) ComputeSize() {
	// abstract
}

func (w Widget) Render(iW IWidget) {
	// abstract
}

func (w *Widget) onMouseIn() {
	// abstract
}

func (w *Widget) onMouseOut() {
	// abstract
}

func (w *Widget) onButtonPress() {
	// abstract
}

func (w *Widget) onButtonRelease() {
	// abstract
}

func (w *Widget) onButtonClick() {
	// abstract
}

func (w *Widget) expand(width, height uint) {
	// abstract
}

func NewWidget(pos geom.Position, w, h uint) *Widget {
	return &Widget{
		pos:          pos,
		w:            w,
		h:            h,
		mouseIn:      false,
		mouseL:       false,
		tip:          "",
		visible:      true,
		disabled:     false,
		back:         rl.Color{40, 40, 120, 255},
		fore:         rl.Color{220, 220, 180, 255},
		backFocus:    rl.Color{70, 70, 130, 255},
		foreFocus:    rl.Color{255, 255, 255, 255},
		backDisabled: rl.Color{40, 40, 120, 125},
		foreDisabled: rl.Color{220, 220, 180, 125},
	}
}
