package gui

import (
	"github.com/go-rogue/engine/cardinal"
	"unicode/utf8"
)

type Button struct {
	Widget
	pressed     bool
	label       string
	BorderStyle BorderStyle
	callback    WidgetCallback
}

func (g *Gui) NewButton(pos cardinal.Position, width, height uint, label string, tip string, borderStyle BorderStyle, callback WidgetCallback, userData interface{}) *Button {
	btn := &Button{
		Widget:      *NewWidget(pos, width, height),
		label:       label,
		callback:    callback,
		BorderStyle: borderStyle,
	}
	btn.tip = tip
	btn.userData = userData
	g.Register(btn)

	return btn
}

func (b *Button) Render(iB IWidget) {
	con := b.gui.con
	fore, back := iB.GetCurrentColors()
	con.SetDefaultForeground(fore)
	con.SetDefaultBackground(back)

	//con.Print(cardinal.Position{X: b.pos.X + int(b.w) / 2, Y: b.pos.Y}, b.label)
	//con.PrintEx(b.x+b.w/2, b.y, BKGND_NONE, CENTER, b.label)

	if b.w > 0 && b.h > 0 {
		con.PrintRectStyle(b.pos, b.w, b.h, SingleWallBorder)
	}
	if b.label != "" {
		con.Print(cardinal.Position{X: b.pos.X + 1, Y: b.pos.Y + 1}, b.label)
		//if b.pressed && b.mouseIn {
		//	con.Print(cardinal.Position{X: b.pos.X + int(b.w)/2, Y: b.pos.Y / 2}, b.label)
		//} else {
		//	con.Print(cardinal.Position{X: b.pos.X + int(b.w)/2, Y: b.pos.Y}, b.label)
		//}
	}
}

func (b *Button) SetLabel(newLabel string) {
	b.label = newLabel
}

func (b *Button) IsPressed() bool {
	return b.pressed
}

func (b *Button) ComputeSize() {
	b.w = uint(utf8.RuneCountInString(b.label)) + 4
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
