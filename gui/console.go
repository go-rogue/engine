package gui

import "C"
import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-rogue/engine/cardinal"
)

type Cell struct {
	char   uint
	fg, bg rl.Color
}

type CellMap map[cardinal.Position]*Cell

type IConsole interface {
	GetData() *CellMap
	GetDefaultBackground() rl.Color
	GetDefaultForeground() rl.Color
	SetDefaultForeground(color rl.Color)
	SetDefaultBackground(color rl.Color)
	Clear()
	GetCharBackground(pos cardinal.Position) rl.Color
	GetCharForeground(pos cardinal.Position) rl.Color
	SetCharBackground(pos cardinal.Position, color rl.Color)
	SetCharForeground(pos cardinal.Position, color rl.Color)
	SetChar(r uint, p cardinal.Position, fg, bg rl.Color)
	//PutChar(x, y, c int, flag BkgndFlag)
	//PutCharEx(x, y, c int, fore, back Color)
	Print(pos cardinal.Position, fmts string, v ...interface{})
	//PrintEx(pos cardinal.Position, flag BkgndFlag, alignment Alignment, fmts string, v ...interface{})
	//PrintRect(x, y, w, h int, fmts string, v ...interface{}) int
	//PrintRectEx(x, y, w, h int, flag BkgndFlag, alignment Alignment, fmts string, v ...interface{}) int
	//HeightRect(x, y, w, h int, fmts string, v ...interface{}) int
	//SetBackgroundFlag(flag BkgndFlag)
	//GetBackgroundFlag() BkgndFlag
	//SetAlignment(alignment Alignment)
	//GetAlignment() Alignment
	//Rect(x, y, w, h int, clear bool, flag BkgndFlag)
	//Hline(x, y, l int, flag BkgndFlag)
	//Vline(x, y, l int, flag BkgndFlag)
	//PrintFrame(x, y, w, h int, empty bool, flag BkgndFlag, fmts string, v ...interface{})
	GetChar(pos cardinal.Position) uint
	GetWidth() uint
	GetHeight() uint
	SetKeyColor(color rl.Color)
	//Blit(xSrc, ySrc, wSrc, hSrc int, dst IConsole, xDst, yDst int, foregroundAlpha, backgroundAlpha float32)
}

type Console struct {
	width, height        uint
	data                 *CellMap
	defaultBg, defaultFg rl.Color
}

// Initiates the CellMap
func (c *Console) init() {
	cMap := make(CellMap)
	c.data = &cMap
}

func (c Console) GetData() *CellMap {
	return c.data
}

func (c Console) GetDefaultBackground() rl.Color {
	return c.defaultBg
}

func (c Console) GetDefaultForeground() rl.Color {
	return c.defaultFg
}

func (c *Console) SetDefaultForeground(color rl.Color) {

}

func (c *Console) SetDefaultBackground(color rl.Color) {

}

func (c *Console) Clear() {

}

func (c Console) GetCharBackground(pos cardinal.Position) rl.Color {
	return (*c.data)[pos].bg
}
func (c Console) GetCharForeground(pos cardinal.Position) rl.Color {
	return (*c.data)[pos].fg
}
func (c *Console) SetCharBackground(pos cardinal.Position, color rl.Color) {
	(*c.data)[pos].bg = color
}
func (c *Console) SetCharForeground(pos cardinal.Position, color rl.Color) {
	(*c.data)[pos].fg = color
}
func (c *Console) SetChar(r uint, p cardinal.Position, fg, bg rl.Color) {
	(*c.data)[p] = &Cell{char: r, fg: fg, bg: bg}
}

func (c *Console) Print(pos cardinal.Position, format string, v ...interface{}) {
	// s := fmt.Sprintf(format, v...)
}

func (c Console) GetChar(pos cardinal.Position) uint {
	return (*c.data)[pos].char
}

func (c Console) GetWidth() uint {
	return c.width
}

func (c Console) GetHeight() uint {
	return c.height
}

func (c *Console) SetKeyColor(color rl.Color) {

}
