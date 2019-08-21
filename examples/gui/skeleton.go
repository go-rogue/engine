package main

import (
	"fmt"
	"github.com/go-rogue/engine/geom"
	"github.com/go-rogue/engine/gui"
	"github.com/go-rogue/engine/scenes"
)

type SkeletonScene struct {
	scenes.Scene
	gui     *gui.Gui
	widgets []gui.IWidget
}

func NewSkeletonScene(gui *gui.Gui) *SkeletonScene {
	return &SkeletonScene{Scene: scenes.Scene{}, gui: gui}
}

func (b *SkeletonScene) Pushed(director *scenes.Director) error {
	b.Director = director
	vBox := b.gui.NewVBox(geom.Point{2, 0}, 0)
	vBox.AddWidget(b.gui.NewButton(geom.Point{0, 0}, 21, 7, "Click to Disable", "", gui.SingleWallBorder, func(w *gui.Button, userData interface{}) {
		w.SetDisabled(true)
		w.SetUserData(w.GetUserData().(int) + 1)
		w.SetLabel(fmt.Sprintf("Clicked %d", w.GetUserData().(int)))
	}, 0))

	vBox.AddWidget(b.gui.NewButton(geom.Point{0, 0}, 21, 7, "Click Me", "", gui.SingleWallBorder, func(w *gui.Button, userData interface{}) {
		x := w.GetUserData().(int)
		if uint(x)+w.GetWidth() <= b.gui.GetConsole().GetWidth() {
			w.SetUserData(w.GetUserData().(int) + 1)
			vBox.SetX(x)
		}
		w.SetLabel(fmt.Sprintf("Clicked %d", x))
	}, 0))

	b.widgets = append(b.widgets, vBox)

	return nil
}

func (b *SkeletonScene) Popped(director *scenes.Director) error {
	for _, w := range b.widgets {
		w.Delete()
	}
	return nil
}
