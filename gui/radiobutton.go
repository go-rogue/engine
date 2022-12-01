package gui

import "github.com/go-rogue/engine/geom"

type RadioButtonGroups struct {
	defaultGroup uint
	groupSelect  [512]*RadioButton
}

func NewRadioButtonGroups() *RadioButtonGroups {
	return &RadioButtonGroups{
		defaultGroup: 0,
		groupSelect:  [512]*RadioButton{},
	}
}

func (g *RadioButtonGroups) SetGroupSelection(btn *RadioButton) {
	g.groupSelect[btn.group] = btn
}

func (g *RadioButtonGroups) UnSelectGroup(group uint) {
	g.groupSelect[group] = nil
}

func (g *RadioButtonGroups) SetDefaultGroup(group uint) {
	g.defaultGroup = group
}

type RadioButtonCallback func(w *RadioButton, userData interface{})

type RadioButton struct {
	Button
	group    uint
	callback RadioButtonCallback
}

func (g *Gui) NewRadioButton(pos geom.Point, width, height uint, label string, tip string, callback RadioButtonCallback, userData interface{}) *RadioButton {
	rBtn := &RadioButton{}
	rBtn.Widget.init(pos, width, height)
	rBtn.label = label
	rBtn.tip = tip
	rBtn.callback = callback
	rBtn.userData = userData
	rBtn.align = BtnTextLeft
	rBtn.labelXPadding = 4
	g.Register(rBtn)
	return rBtn
}

func (b *RadioButton) Select() {
	b.gui.rbs.SetGroupSelection(b)
}

func (b *RadioButton) UnSelect() {
	b.gui.rbs.UnSelectGroup(b.group)
}

func (b *RadioButton) IsSelected() bool {
	return b.gui.rbs.groupSelect[b.group] == b
}

func (b *RadioButton) SetGroup(group uint) {
	b.group = group
}

func (b *RadioButton) Render(iB IWidget) {
	con := b.gui.con
	fore, back := iB.GetCurrentColors()
	con.SetDefaultForeground(fore)
	con.SetDefaultBackground(back)

	b.Button.Render(iB)

	if b.IsSelected() {
		var labelPos = b.GetLabelPos("> " + b.label)

		// Radio Buttons have an X padding of 4 cells, this gives one cell padding each side of the selected
		// marker. Here we pull the x pos back two cells so we have a space between the selected marker and
		// the label.
		labelPos.X -= 2
		con.PutCharEx('>', labelPos, fore, back)
	}
}

func (b *RadioButton) onButtonClick() {
	if b.disabled {
		return
	}

	b.Select()
	b.Button.onButtonClick()
}