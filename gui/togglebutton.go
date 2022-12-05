package gui

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
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
	btn.label = FrameTitle{
		text:               label,
		alignment:          AlignTextCenter,
		verticallyCentered: true,
		padding:            rl.NewVector2(2, 0),
	}
	btn.callback = callback
	btn.tip = tip
	btn.userData = userData
	btn.pressed = false

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
	var labelPos = b.label.Position(b.pos, b.w, b.h)
	var label string
	var icon rune

	con := b.gui.con
	fore, back := iB.GetCurrentColors()
	con.SetDefaultForeground(fore)
	con.SetDefaultBackground(back)

	b.Button.Render(iB)

	if b.toggled {
		icon = sprites.TCOD_CHAR_CHECKBOX_SET
	} else {
		icon = sprites.TCOD_CHAR_CHECKBOX_UNSET
	}

	if b.label.IsVisible() {
		label = fmt.Sprintf("%c %s", icon, b.label.text)
	} else {
		label = fmt.Sprintf("%c", icon)
	}

	con.Print(labelPos, label)
}

func (b *ToggleButton) onButtonClick() {
	if b.disabled {
		return
	}

	b.toggled = !b.toggled
	b.Button.onButtonClick()
}
