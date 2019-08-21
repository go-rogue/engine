package main

import (
	"github.com/go-rogue/engine/gui"
	"github.com/go-rogue/engine/scenes"
)

type SkeletonScene struct {
	scenes.Scene
	gui     *gui.Gui
	widgets []gui.IWidget
}

func NewSkeletonScene(gui *gui.Gui) *SkeletonScene {
	return &SkeletonScene{Scene: scenes.Scene{Name: "Skeleton"}, gui: gui}
}

func (b *SkeletonScene) Pushed(director *scenes.Director) error {
	b.Director = director
	b.widgets = make([]gui.IWidget, 0)

	return nil
}

func (b *SkeletonScene) Popped(director *scenes.Director) error {
	for _, w := range b.widgets {
		w.Delete()
	}
	return nil
}
