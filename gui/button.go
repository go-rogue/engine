package gui

import (
	"github.com/go-rogue/engine/geom"
	"unicode/utf8"
)

type btnTextAlign uint

type ButtonCallback func(w *Button, userData interface{})

const (
	BtnTextLeft btnTextAlign = iota
	BtnTextCenter
	BtnTextRight
)

type Button struct {
	Widget
	pressed       bool
	label         string
	labelXPadding int
	BorderStyle   FrameStyle
	callback      ButtonCallback
	align         btnTextAlign
}

func (g *Gui) NewButton(pos geom.Point, width, height uint, label string, tip string, borderStyle FrameStyle, callback ButtonCallback, userData interface{}) *Button {
	btn := &Button{}
	btn.Widget.init(pos, width, height)
	btn.label = label
	btn.callback = callback
	btn.BorderStyle = borderStyle
	btn.align = BtnTextCenter
	btn.tip = tip
	btn.userData = userData
	btn.labelXPadding = 2
	g.Register(btn)
	return btn
}

func (b *Button) SetTextAlign(a btnTextAlign) {
	b.align = a
}

func (b *Button) Render(iB IWidget) {
	con := b.gui.con
	fore, back := iB.GetCurrentColors()
	con.SetDefaultForeground(fore)
	con.SetDefaultBackground(back)

	if b.w > 0 && b.h > 0 {
		con.PrintFrame(b.pos, b.w, b.h, b.BorderStyle, true, true)
	}
	if b.label != "" {
		con.Print(b.GetLabelPos(b.label), b.label)
	}
}

func (b Button) GetLabelPos(label string) geom.Point {
	var txtPos geom.Point

	if b.align == BtnTextLeft {
		txtPos = geom.Point{X: b.pos.X + b.labelXPadding, Y: b.pos.Y + int(b.h/2)}
	} else if b.align == BtnTextCenter {
		txtPos = geom.Point{X: b.pos.X + int(int(b.w/2)-utf8.RuneCountInString(label)/2), Y: b.pos.Y + int(b.h/2)}
	} else if b.align == BtnTextRight {
		txtPos = geom.Point{X: b.pos.X + int(int(b.w)-utf8.RuneCountInString(label)-b.labelXPadding), Y: b.pos.Y + int(b.h/2)}
	}

	return txtPos
}

func (b *Button) SetLabel(newLabel string) {
	b.label = newLabel
}

func (b Button) GetLabel() string {
	return b.label
}

func (b *Button) IsPressed() bool {
	return b.pressed
}

func (b *Button) ComputeSize() {
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
