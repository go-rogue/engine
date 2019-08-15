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

func NewGui(console IConsole) *Gui {
	return &Gui{
		con:          console,
		widgetVector: []IWidget{},
	}
}
