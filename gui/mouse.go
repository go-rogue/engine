package gui

import "github.com/go-rogue/engine/geom"

var MouseStatus = &Mouse{}

var UpdateMouseStatus func()

type Mouse struct {
	Pos            geom.Position
	LButton        bool
	RButton        bool
	MButton        bool
	LButtonPressed bool
	RButtonPressed bool
	MButtonPressed bool
	WheelUp        bool
	WheelDown      bool
	WindowFocus    bool
	Visible        bool
	Supported      bool
}
