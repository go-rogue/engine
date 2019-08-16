package gui

import "C"
import (
	"bytes"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-rogue/engine/cardinal"
)

type CellLayer struct {
	char   uint
	fg, bg rl.Color
}

type Cell struct {
	char   uint
	fg, bg rl.Color
}

type CellMap map[cardinal.Position]*Cell

type IConsole interface {
	GetData() *CellMap
	GetDefaultBackground() rl.Color
	GetDefaultForeground() rl.Color
	SetDefaultForeground(colour rl.Color)
	SetDefaultBackground(colour rl.Color)
	Clear()
	GetCharBackground(pos cardinal.Position) rl.Color
	GetCharForeground(pos cardinal.Position) rl.Color
	SetCharBackground(pos cardinal.Position, colour rl.Color)
	SetCharForeground(pos cardinal.Position, colour rl.Color)
	SetChar(r uint, p cardinal.Position, fg, bg rl.Color)
	PutChar(r uint, p cardinal.Position)
	//PutCharEx(x, y, c int, fore, back Color)
	Print(pos cardinal.Position, str string)
	SPrintf(pos cardinal.Position, fmts string, v ...interface{})
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
	SetKeyColor(colour rl.Color)
	Blit(xSrc, ySrc, wSrc, hSrc int, dst IConsole, xDst, yDst int, foregroundAlpha, backgroundAlpha float32)
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

func (c *Console) SetDefaultForeground(colour rl.Color) {
	c.defaultFg = colour
}

func (c *Console) SetDefaultBackground(colour rl.Color) {
	c.defaultBg = colour
}

//
// Clear the console data
//
func (c *Console) Clear() {
	c.init()
}

func (c Console) GetCharBackground(pos cardinal.Position) rl.Color {
	return (*c.data)[pos].bg
}
func (c Console) GetCharForeground(pos cardinal.Position) rl.Color {
	return (*c.data)[pos].fg
}
func (c *Console) SetCharBackground(pos cardinal.Position, colour rl.Color) {
	(*c.data)[pos].bg = colour
}
func (c *Console) SetCharForeground(pos cardinal.Position, colour rl.Color) {
	(*c.data)[pos].fg = colour
}

func (c *Console) PutChar(r uint, p cardinal.Position) {
	(*c.data)[p] = &Cell{char: r, fg: c.defaultFg, bg: c.defaultBg}
}

func (c *Console) SetChar(r uint, p cardinal.Position, fg, bg rl.Color) {
	(*c.data)[p] = &Cell{char: r, fg: fg, bg: bg}
}

func (c *Console) Print(pos cardinal.Position, str string) {
	//strLen := uint(pos.X + utf8.RuneCountInString(str))

	// Split str into individual runes and then loop over and set the
	// char at that cardinal position.

	// <%FG:colour_name>Text<%/>
	// <%BG:colour_name>Text<%/>
	// <%FG:colour_name,BG:colour_name>Text<%/>
	// <%RESET/>

	split := []rune(str)
	fg := c.defaultFg
	bg := c.defaultBg
	readCommand := false
	buf := bytes.Buffer{}

	xOff := 0 // Remove x spacing added by command characters

	for i, r := range split {
		if (r == '<' && split[i+1] == '%' && (split[i+2] == 'F' || split[i+2] == 'B' || split[i+2] == 'R')) || (r == '<' && split[i+1] == '%') {
			readCommand = true
			xOff--
			continue
		} else if r == '>' && readCommand == true {
			readCommand = false

			command := buf.String()
			if command != "" {
				fmt.Println("cmd: ", command)
			}
			buf.Reset()
			xOff--
			continue
		}

		if readCommand {
			xOff--
			if r != '%' && r != '/' {
				buf.WriteRune(r)
			}
		} else {
			(*c.data)[cardinal.Position{X: pos.X + i + xOff, Y: pos.Y}] = &Cell{char: uint(r), bg: bg, fg: fg}
		}
	}
}

// May not actually use this...
func (c *Console) SPrintf(pos cardinal.Position, format string, v ...interface{}) {
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

func (c *Console) SetKeyColor(colour rl.Color) {

}

//
// This function allows you to blit a rectangular area of the source console at a specific position on a destination
// console. It can also simulate alpha transparency with the fade parameter.
//
func (c *Console) Blit(xSrc, ySrc, wSrc, hSrc int, dst IConsole, xDst, yDst int, foregroundAlpha, backgroundAlpha float32) {
	// @todo
}

//
// You can create as many off-screen consoles as you want by using this function. Then use the
// blit function to draw content to a windowed virtual console such as RaylibConsole.
//
func NewVirtualConsole(w, h uint) *Console {
	ret := &Console{width: w, height: h, defaultBg: rl.Black, defaultFg: rl.White}
	ret.init()
	return ret
}
