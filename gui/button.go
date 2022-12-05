package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-rogue/engine/geom"
)

type ButtonCallback func(w *Button, userData interface{})

type Button struct {
	Widget
	pressed  bool
	label    FrameTitle
	callback ButtonCallback
}

func (g *Gui) NewButton(pos geom.Point, width, height uint, label string, tip string, borderStyle FrameStyle, callback ButtonCallback, userData interface{}) *Button {
	btn := &Button{}
	btn.Widget.init(pos, width, height, borderStyle)
	btn.label = FrameTitle{
		text:               label,
		alignment:          AlignTextCenter,
		verticallyCentered: true,
		padding:            rl.NewVector2(2, 0),
	}
	btn.callback = callback
	btn.tip = tip
	btn.userData = userData
	g.Register(btn)
	return btn
}

func (g *Gui) NewBasicButton(pos geom.Point, width, height uint, label string, borderStyle FrameStyle) *Button {
	return g.NewButton(pos, width, height, label, "", borderStyle, func(w *Button, userData interface{}) {}, nil)
}

func (b *Button) Render(iB IWidget) {
	con := b.gui.con
	fore, back := iB.GetCurrentColors()
	con.SetDefaultForeground(fore)
	con.SetDefaultBackground(back)

	if b.w > 0 && b.h > 0 {
		con.PrintFrame(b.pos, b.w, b.h, b.borderStyle, ZeroFrameTitle, true, true)
	}
	if b.label.IsVisible() {
		con.Print(b.label.Position(b.pos, b.w, b.h), b.label.text)
	}
}

func (b *Button) SetLabel(newLabel string) {
	b.label.text = newLabel
}

func (b Button) GetLabel() string {
	return b.label.text
}

func (b *Button) IsPressed() bool {
	return b.pressed
}

func (b *Button) ComputeSize() {
	// TODO... is this used?
	// b.w = uint(utf8.RuneCountInString(b.label)) + 4
}

func (b *Button) onButtonPress() {
	b.pressed = true
}

func (b *Button) onButtonRelease() {
	b.pressed = false
}

func (b *Button) onButtonClick() {
	if b.disabled == false && b.callback != nil {
		b.callback(b, b.userData)
	}
}

func (b *Button) expand(width, height uint) {
	if b.w < width {
		b.w = width
	}
}
