package gui

import "github.com/go-rogue/engine/cardinal"

var MouseStatus = &Mouse{}

var UpdateMouseStatus func()

type Mouse struct {
	Pos            cardinal.Position
	LButton        bool
	RButton        bool
	MButton        bool
	LButtonPressed bool
	RButtonPressed bool
	MButtonPressed bool
	WheelUp        bool
	WheelDown      bool
	WindowFocus    bool
}
