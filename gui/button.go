package gui

import (
	"github.com/go-rogue/engine/cardinal"
	"unicode/utf8"
)

type btnTextAlign uint

type ButtonCallback func(w *Button, userData interface{})

const (
	BtnTextLeft btnTextAlign = iota
	BtnTextCenter
	BtnTextRight
)

// @todo add disabled flag
type Button struct {
	Widget
	pressed     bool
	label       string
	BorderStyle BorderStyle
	callback    ButtonCallback
	align       btnTextAlign
}

func (g *Gui) NewButton(pos cardinal.Position, width, height uint, label string, tip string, borderStyle BorderStyle, callback ButtonCallback, userData interface{}) *Button {
	btn := &Button{
		Widget:      *NewWidget(pos, width, height),
		label:       label,
		callback:    callback,
		BorderStyle: borderStyle,
		align:       BtnTextCenter,
	}
	btn.tip = tip
	btn.userData = userData
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
		con.PrintRectStyle(b.pos, b.w, b.h, SingleWallBorder, true, true)
	}
	if b.label != "" {
		var txtPos cardinal.Position
		padX := 2

		if b.align == BtnTextLeft {
			txtPos = cardinal.Position{b.pos.X + int(padX), b.pos.Y + int(b.h/2)}
		} else if b.align == BtnTextCenter {
			txtPos = cardinal.Position{b.pos.X + int(int(b.w/2)-utf8.RuneCountInString(b.label)/2), b.pos.Y + int(b.h/2)}
		} else if b.align == BtnTextRight {
			txtPos = cardinal.Position{b.pos.X + int(int(b.w)-utf8.RuneCountInString(b.label)-int(padX)), b.pos.Y + int(b.h/2)}
		}

		con.Print(txtPos, b.label)
	}
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
	if b.callback != nil {
		b.callback(b, b.userData)
	}
}

func (b *Button) expand(width, height uint) {
	if b.w < width {
		b.w = width
	}
}
