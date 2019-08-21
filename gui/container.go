package gui

import "github.com/go-rogue/engine/geom"

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
		c.gui.Unregister(w)
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
