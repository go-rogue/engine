package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-rogue/engine/cardinal"
)

type IWidget interface {
	SetGui(*Gui)
	GetGui() *Gui
	Delete()
	GetWidth() uint
	SetWidth(width uint)
	GetHeight() uint
	SetHeight(height uint)
	GetPosition() cardinal.Position
	Move(pos cardinal.Position)
	IsVisible() bool
	SetVisible(visible bool)
	GetUserData() interface{}
	SetUserData(data interface{})
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
	expand(x, y int)
}

type Widget struct {
	pos       cardinal.Position
	w, h      uint
	userData  interface{}
	tip       string
	mouseIn   bool
	mouseL    bool
	visible   bool
	back      rl.Color
	fore      rl.Color
	backFocus rl.Color
	foreFocus rl.Color
	gui       *Gui
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

func (w Widget) GetPosition() cardinal.Position {
	return w.pos
}

func (w *Widget) Move(pos cardinal.Position) {
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
	if w.mouseIn {
		fore = w.foreFocus
		back = w.backFocus
	} else {
		fore = w.fore
		back = w.back
	}

	return fore, back
}

func (w *Widget) Update(iW IWidget) {
	/// ...
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

func (w *Widget) expand(x, y int) {
	// abstract
}

func NewWidget(pos cardinal.Position, w, h uint) *Widget {
	return &Widget{
		pos:       pos,
		w:         w,
		h:         h,
		mouseIn:   false,
		mouseL:    false,
		tip:       "",
		visible:   true,
		back:      rl.Color{40, 40, 120, 255},
		fore:      rl.Color{220, 220, 180, 255},
		backFocus: rl.Color{70, 70, 130, 255},
		foreFocus: rl.Color{255, 255, 255, 255},
	}
}
