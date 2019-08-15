package gui

import "github.com/go-rogue/engine/cardinal"

type Mouse struct {
	hover bool
	x, y  int
	pos   cardinal.Position
}
