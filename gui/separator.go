package gui

import (
	"github.com/go-rogue/engine/geom"
	"github.com/go-rogue/engine/sprites"
)

type Separator struct {
	Widget
	text string
}

func (g *Gui) NewSeparator(text, tip string) *Separator {
	separator := &Separator{}
	separator.init(geom.Point{X: 0, Y: 0}, 0, 0, ZeroWallBorder)
	separator.text = text
	separator.tip = tip

	g.Register(separator)
	return separator
}

func (s *Separator) ComputeSize() {
	if s.text != "" {
		s.w = uint(len(s.text) + 2)
	} else {
		s.w = 0
	}
}

func (s *Separator) expand(width, height uint) {
	if s.w < width {
		s.w = width
	}
}

func (s *Separator) Render(iW IWidget) {
	con := s.gui.con
	con.SetDefaultBackground(s.back)
	con.SetDefaultForeground(s.fore)

	con.Hline(s.pos, s.w)

	if con.InBounds(geom.Point{X: s.pos.X - 1, Y: s.pos.Y}) {
		con.SetChar(sprites.TCOD_CHAR_TEEE, geom.Point{X: s.pos.X - 1, Y: s.pos.Y}) // ├
	}

	if con.InBounds(geom.Point{X: s.pos.X + int(s.w), Y: s.pos.Y}) {
		con.SetChar(sprites.TCOD_CHAR_TEEW, geom.Point{X: s.pos.X + int(s.w), Y: s.pos.Y}) // ┤
	}

	con.SetDefaultBackground(s.fore)
	con.SetDefaultForeground(s.back)

	// TODO: need to port con.PrintEx
	// con.PrintEx(self.x+self.w/2, self.y, BKGND_SET, CENTER, " %s ", self.txt)
}
