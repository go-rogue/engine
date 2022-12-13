package gui

import "github.com/go-rogue/engine/geom"

type Toolbar struct {
	Container
	fixedWidth uint
	name       FrameTitle
}

func (g *Gui) NewToolbar(pos geom.Point, width uint, name string) *Toolbar {
	toolbar := &Toolbar{
		name: FrameTitle{
			text:               name,
			alignment:          AlignTextCenter,
			verticallyCentered: false,
		},
	}

	toolbar.init(pos, width, 2)
	toolbar.borderStyle = SingleWallBorder

	if width == 0 {
		toolbar.w = uint(len(name) + 4)
		toolbar.fixedWidth = 0
	} else if uint(len(name)+4) > width {
		toolbar.w = uint(len(name) + 4)
		toolbar.fixedWidth = uint(len(name) + 4)
	} else {
		toolbar.w = width
		toolbar.fixedWidth = width
	}

	g.Register(toolbar)
	return toolbar
}

func (t *Toolbar) SetName(name string) {
	t.name.text = name

	nFixedWidth := uint(len(t.name.text) + 4)
	if nFixedWidth > t.fixedWidth {
		t.fixedWidth = nFixedWidth
	}
}

func (t *Toolbar) AddSeparator(txt, tip string) {
	t.AddWidget(t.gui.NewSeparator(txt, tip))
}

func (t *Toolbar) ComputeSize() {
	cY := t.pos.Y + 1
	if t.name.text != "" {
		t.w = uint(len(t.name.text) + 4)
	} else {
		t.w = 2
	}

	for _, c := range t.content {
		if c.IsVisible() {
			c.SetX(t.pos.X + 1)
			c.SetY(cY)
			c.ComputeSize()
			if c.GetWidth()+2 > t.w {
				t.w = c.GetWidth() + 2
			}
			cY += int(t.GetHeight())
		}
	}

	if t.w < t.fixedWidth {
		t.w = t.fixedWidth
	}

	// t.h = uint(cY - t.pos.Y + 1)

	for _, c := range t.content {
		if c.IsVisible() {
			c.expand(t.w-2, c.GetHeight())
		}
	}
}

func (t *Toolbar) Render(iW IWidget) {
	con := t.gui.con
	con.SetDefaultBackground(t.back)
	con.SetDefaultForeground(t.fore)

	// TODO, I am going to apply a frame here, but this should exist on Container and switched on by default for Toolbar
	con.PrintFrame(t.pos, t.w, t.h, t.borderStyle, t.name, false, true)

	t.Container.Render(iW)
}
