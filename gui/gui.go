package gui

type Gui struct {
	focus         IWidget // focused widget
	keyboardFocus IWidget // keyboard focused widget
	mouse         Mouse
	elapsed       float32
	con           IConsole
	widgetVector  []IWidget
	//rbs           *RadioButtonStatic
	//tbs           *TextBoxStatic
}

//
// Add a Widget to the Gui
//
func (g *Gui) Register(w IWidget) {
	w.SetGui(g)
	g.widgetVector = append(g.widgetVector, w)
}

//
// Remove a Widget from the Gui
//
func (g *Gui) Unregister(w IWidget) {
	if g.focus == w {
		g.focus = nil
	}
	if g.keyboardFocus == w {
		g.keyboardFocus = nil
	}
	for i, e := range g.widgetVector {
		if e == w {
			g.widgetVector = append(g.widgetVector[0:i], g.widgetVector[i+1:]...)
		}
	}
}

func (g *Gui) UpdateWidgets(dt float32) {
	g.mouse = *MouseStatus
	g.elapsed = dt

	for _, w := range g.widgetVector {
		if w.IsVisible() {
			w.ComputeSize()
			w.Update(w)
		}
	}
}

func (g *Gui) RenderWidgets() {
	for _, w := range g.widgetVector {
		if w.IsVisible() {
			fore, back := g.con.GetDefaultForeground(), g.con.GetDefaultBackground()
			w.Render(w)
			g.con.SetDefaultForeground(fore)
			g.con.SetDefaultBackground(back)
		}
	}
}

//
// Set the Console the Gui should render to
//
func (g *Gui) SetConsole(console IConsole) {
	g.con = console
}

func (g *Gui) GetConsole() IConsole {
	return g.con
}

//
// Set the focused Widget for mouse input
//
func (g *Gui) IsFocused(w IWidget) bool {
	return g.focus == w
}

//
// Set the focused Widget for keyboard input
//
func (g *Gui) IsKeyboardFocused(w IWidget) bool {
	return g.keyboardFocus == w
}

//
// Get the focused Widget for mouse input
//
func (g *Gui) GetFocusedWidget() IWidget {
	return g.focus
}

//
// Get the focused Widget for keyboard input
//
func (g *Gui) GetFocusedKeyboardWidget() IWidget {
	return g.keyboardFocus
}

//
// Construct a new Gui
//
func NewGui(console IConsole) *Gui {
	return &Gui{
		con:          console,
		widgetVector: []IWidget{},
	}
}
