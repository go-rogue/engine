package gui

import (
	"github.com/go-rogue/engine/geom"
)

// Container is a workspace for displaying Widgets within a particular "viewport"
type Container struct {
	Widget
	content []IWidget
}

func (g *Gui) NewContainer(pos geom.Point, width, height uint) *Container {
	container := &Container{}
	container.init(pos, width, height)
	g.Register(container)
	return container
}

func (c *Container) init(pos geom.Point, width, height uint) {
	c.Widget.init(pos, width, height)
	c.content = []IWidget{}
}

func (c *Container) AddWidget(w IWidget) {
	c.content = append(c.content, w)
	c.gui.Unregister(w)
}

func (c *Container) RemoveWidget(w IWidget) {
	for i, e := range c.content {
		if e == w {
			c.content = append(c.content[:i], c.content[i+1:]...)
		}
	}
}

func (c *Container) Render(iW IWidget) {

	// Uncomment out the below to see the render box for this container, you will need to comment out w.Render as well
	// in order to see the box

	//for x := uint(c.pos.X); x <= (uint(c.pos.X) + c.w); x++ {
	//	for y := uint(c.pos.Y); y < (uint(c.pos.Y) + c.h); y++ {
	//		c.gui.con.PutCharEx(' ', geom.Point{X:int(x), Y:int(y)}, rl.Color{255, 0, 0, 255}, rl.Color{255, 0, 0, 255})
	//	}
	//}

	for _, w := range c.content {
		if w.IsVisible() {
			w.Render(w)
		}
	}
}

func (c *Container) Clear() {
	c.content = []IWidget{}
}

func (c *Container) Delete() {
	for _, w := range c.content {
		c.RemoveWidget(w)
	}
	c.gui.Unregister(c)
}

func (c *Container) Update(iW IWidget) {
	c.Widget.Update(iW)
	for _, w := range c.content {
		if w.IsVisible() {
			w.Update(w)
		}
	}
}
