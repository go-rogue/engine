package gui

import "github.com/go-rogue/engine/geom"

// HBox is an extension to Container that ensures Widget's added to it are presented in a row
type HBox struct {
	VBox
}

func (g *Gui) NewHBox(pos geom.Point, padding uint) *HBox {
	hBox := &HBox{}
	g.Register(hBox)
	hBox.init(pos, padding)
	return hBox
}

func (h *HBox) ComputeSize() {
	currentX := h.pos.X
	h.h = 0

	for _, w := range h.content {
		if w.IsVisible() {
			w.SetX(currentX)
			w.SetY(h.pos.Y)
			w.ComputeSize()
			if w.GetHeight() > h.h {
				h.h = w.GetHeight()
			}
			currentX += int(w.GetWidth() + h.padding)
		}
	}

	h.w = uint(currentX - h.pos.X)

	for _, w := range h.content {
		if w.IsVisible() {
			w.expand(w.GetWidth(), h.h)
		}
	}
}
