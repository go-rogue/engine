package gui

import "github.com/go-rogue/engine/geom"

type VBox struct {
	Container
	padding uint
}

func (g *Gui) NewVBox(pos geom.Point, padding uint) *VBox {
	vBox := &VBox{}
	g.Register(vBox)
	vBox.init(pos, padding)
	return vBox
}

func (v *VBox) init(pos geom.Point, padding uint) {
	v.Container.init(pos, 0, 0)
	v.padding = padding
}

func (v *VBox) ComputeSize() {
	currentY := v.pos.Y
	v.w = 0
	for _, w := range v.content {
		if w.IsVisible() {
			w.SetX(v.pos.X)
			w.SetY(currentY)
			w.ComputeSize()
			if w.GetWidth() > v.w {
				v.w = w.GetWidth()
			}
			currentY += int(w.GetHeight() + v.padding)
		}
	}
	v.h = uint(currentY - v.pos.Y)

	for _, w := range v.content {
		if w.IsVisible() {
			w.expand(v.w, w.GetHeight())
		}
	}
}
