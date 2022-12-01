package main

import (
	"fmt"
	"github.com/go-rogue/engine/geom"
	"github.com/go-rogue/engine/gui"
	"github.com/go-rogue/engine/scenes"
)

type ButtonScene struct {
	scenes.Scene
	gui     *gui.Gui
	widgets []gui.IWidget
}

func NewButtonScene(gui *gui.Gui) *ButtonScene {
	return &ButtonScene{Scene: scenes.Scene{Name: "VBox Button Example"}, gui: gui}
}

func (b *ButtonScene) Pushed(director *scenes.Director) error {
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

	vBox2 := b.gui.NewVBox(geom.Point{2, 30}, 0)

	rBtn1 := b.gui.NewBasicRadioButton(geom.Point{0, 0}, 21, 7, "Group 1.0", gui.SingleWallBorder)
	rBtn2 := b.gui.NewBasicRadioButton(geom.Point{0, 0}, 21, 7, "Group 1.1", gui.SingleWallBorder)

	rBtn1.SetGroup(0)
	rBtn2.SetGroup(0)

	vBox2.AddWidget(rBtn1)
	vBox2.AddWidget(rBtn2)
	b.widgets = append(b.widgets, rBtn1)
	b.widgets = append(b.widgets, rBtn2)

	hBox := b.gui.NewHBox(geom.Point{2, 20}, 0)

	rBtn3 := b.gui.NewBasicRadioButton(geom.Point{0, 0}, 21, 7, "Group 2.0", gui.SingleWallBorder)
	rBtn4 := b.gui.NewBasicRadioButton(geom.Point{0, 0}, 21, 7, "Group 2.1", gui.SingleWallBorder)
	rBtn5 := b.gui.NewBasicRadioButton(geom.Point{0, 0}, 21, 7, "Group 2.2", gui.SingleWallBorder)

	rBtn3.SetGroup(1)
	rBtn4.SetGroup(1)
	rBtn5.SetGroup(1)

	hBox.AddWidget(rBtn3)
	hBox.AddWidget(rBtn4)
	hBox.AddWidget(rBtn5)

	hBox2 := b.gui.NewHBox(geom.Point{23, 37}, 0)
	tBtn := b.gui.NewBasicToggleButton(geom.Point{0, 0}, 7, 7, "", gui.SingleWallBorder)
	tBtn2 := b.gui.NewBasicToggleButton(geom.Point{0, 0}, 21, 7, "Toggle Btns", gui.SingleWallBorder)
	hBox2.AddWidget(tBtn)
	hBox2.AddWidget(tBtn2)

	b.widgets = append(b.widgets, rBtn3)
	b.widgets = append(b.widgets, rBtn4)
	b.widgets = append(b.widgets, rBtn5)
	b.widgets = append(b.widgets, tBtn)

	return nil
}

func (b *ButtonScene) Popped(director *scenes.Director) error {
	for _, w := range b.widgets {
		w.Delete()
	}
	return nil
}
