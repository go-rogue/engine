package gui

import (
	"fmt"
	"github.com/go-rogue/engine/geom"
	"github.com/go-rogue/engine/sprites"
)

type ToggleButtonCallback func(w *ToggleButton, userData interface{})

type ToggleButton struct {
	Button
	callback ToggleButtonCallback
	toggled  bool
}

func (g *Gui) NewToggleButton(pos geom.Point, width, height uint, label string, tip string, borderStyle FrameStyle, callback ToggleButtonCallback, userData interface{}) *ToggleButton {
	btn := &ToggleButton{}
	btn.Widget.init(pos, width, height, borderStyle)
	btn.label = label
	btn.callback = callback
	btn.align = BtnTextCenter
	btn.tip = tip
	btn.userData = userData
	btn.pressed = false
	btn.labelXPadding = 2

	g.Register(btn)
	return btn
}

func (g *Gui) NewBasicToggleButton(pos geom.Point, width, height uint, label string, borderStyle FrameStyle) *ToggleButton {
	return g.NewToggleButton(pos, width, height, label, "", borderStyle, func(w *ToggleButton, userData interface{}) {}, nil)
}

func (b ToggleButton) IsToggled() bool {
	return b.toggled
}

func (b *ToggleButton) SetToggled(value bool) {
	b.toggled = value
}

func (b *ToggleButton) Render(iB IWidget) {
	con := b.gui.con
	fore, back := iB.GetCurrentColors()
	con.SetDefaultForeground(fore)
	con.SetDefaultBackground(back)

	b.Button.Render(iB)

	var icon rune
	if b.toggled {
		icon = sprites.TCOD_CHAR_CHECKBOX_SET
	} else {
		icon = sprites.TCOD_CHAR_CHECKBOX_UNSET
	}

	if b.label == "" {
		label := fmt.Sprintf("%c", icon)
		con.Print(b.GetLabelPos(label), label)
	} else {
		label := fmt.Sprintf("%c %s", icon, b.label)
		con.Print(b.GetLabelPos(label), label)
	}
}

func (b *ToggleButton) onButtonClick() {
	if b.disabled {
		return
	}

	b.toggled = !b.toggled
	b.Button.onButtonClick()
}
