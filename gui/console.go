package gui

import "C"
import (
	"bytes"
	"container/list"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-rogue/engine/cardinal"
	"strings"
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

type printCommand struct {
	fg rl.Color
	bg rl.Color
}

func (p printCommand) AsFgBg() (rl.Color, rl.Color) {
	return p.fg, p.bg
}

func (p printCommand) Copy() *printCommand {
	return &printCommand{fg: p.fg, bg: p.bg}
}

type printVM struct {
	queue *list.List
}

// tmp
// @todo replace with theme service
func StrToColour(str string) rl.Color {
	if str == "white" {
		return rl.White
	} else if str == "black" {
		return rl.Black
	} else if str == "red" {
		return rl.Red
	} else if str == "blue" {
		return rl.Blue
	}

	return rl.Black
}

func (p printVM) iExecute(cmd string, cCmd *printCommand) {
	command := strings.Split(cmd, ":")
	which := strings.ToLower(command[0])
	colour := StrToColour(command[1])

	if which == "fg" {
		cCmd.fg = colour
	} else if which == "bg" {
		cCmd.bg = colour
	}
}

func (p *printVM) Execute(cmd string) {
	if cmd == "" {
		return
	}

	cCmd := p.Peek().Copy()

	if strings.ContainsRune(cmd, ',') {
		commands := strings.Split(cmd, ",")
		for _, c := range commands {
			p.iExecute(c, cCmd)
		}
	} else {
		p.iExecute(cmd, cCmd)
	}

	p.queue.PushBack(*cCmd)
}

func (p *printVM) Peek() printCommand {
	return p.queue.Back().Value.(printCommand)
}

func (p *printVM) Remove() {
	if p.queue.Len() == 1 {
		return
	}
	p.queue.Remove(p.queue.Back())
}

func newPrintVM(initial printCommand) *printVM {
	ret := &printVM{queue: list.New()}
	ret.queue.PushBack(initial)
	return ret
}

type IConsole interface {
	GetData() *CellMap
	GetCellAtPos(pos cardinal.Position) *Cell
	GetDefaultBackground() rl.Color
	GetDefaultForeground() rl.Color
	SetDefaultForeground(colour rl.Color)
	SetDefaultBackground(colour rl.Color)
	Clear()
	GetCharBackground(pos cardinal.Position) rl.Color
	GetCharForeground(pos cardinal.Position) rl.Color
	SetCharBackground(pos cardinal.Position, colour rl.Color)
	SetCharForeground(pos cardinal.Position, colour rl.Color)
	SetChar(r uint, pos cardinal.Position)
	PutChar(r uint, p cardinal.Position)
	PutCharEx(r uint, p cardinal.Position, fg, bg rl.Color)
	Print(pos cardinal.Position, str string)
	//PrintRect(x, y, w, h int, fmts string, v ...interface{}) int
	ClearRect(pos cardinal.Position, w, h uint)
	PrintRectStyle(pos cardinal.Position, w, h uint, boxStyle BorderStyle, filled, clear bool)
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

	for cY := 0; cY < int(c.height); cY++ {
		for cX := 0; cX < int(c.width); cX++ {
			c.PutChar(' ', cardinal.Position{X: cX, Y: cY})
		}
	}
}

func (c Console) GetData() *CellMap {
	return c.data
}

func (c *Console) GetCellAtPos(pos cardinal.Position) *Cell {
	return (*c.data)[pos]
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

func (c *Console) SetChar(r uint, pos cardinal.Position) {
	(*c.data)[pos].char = r
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

func (c *Console) PutCharEx(r uint, p cardinal.Position, fg, bg rl.Color) {
	(*c.data)[p] = &Cell{char: r, fg: fg, bg: bg}
}

//
// Split str into individual runes and then loop over and set the
// char at that cardinal position.
//
// <%FG:colour_name>Text<%/>
// <%BG:colour_name>Text<%/>
// <%FG:colour_name,BG:colour_name>Text<%/>
//
func (c *Console) Print(pos cardinal.Position, str string) {
	// Split string into runes
	split := []rune(str)

	vm := newPrintVM(printCommand{fg: c.defaultFg, bg: c.defaultBg})

	// default command read to false
	readCommand := false

	// Setup string buffer
	buf := bytes.Buffer{}

	// Remove x spacing added by command characters
	xOff := 0

	for i, r := range split {
		if r == '<' && split[i+1] == '%' && (split[i+2] == 'F' || split[i+2] == 'B') {
			// Begin opening command <%...>
			readCommand = true
			xOff--
			continue
		} else if r == '<' && split[i+1] == '%' && split[i+2] == '/' && split[i+3] == '>' {
			// Begin closing command <%/>
			vm.Remove()
			readCommand = true
			xOff--
			continue
		} else if r == '>' && readCommand == true {
			// Complete reading command, if this is not a <%/> then
			// the string passed to vm will not be empty.
			readCommand = false
			vm.Execute(buf.String())
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
			fg, bg := vm.Peek().AsFgBg()
			(*c.data)[cardinal.Position{X: pos.X + i + xOff, Y: pos.Y}] = &Cell{char: uint(r), bg: bg, fg: fg}
		}
	}
}

//
// Reset cells within the Rectangular area to default
//
func (c *Console) ClearRect(pos cardinal.Position, w, h uint) {
	fmt.Println(pos, w, h)

	for y := 0; y < int(h)-1; y++ {
		for x := 0; x < int(w); x++ {
			c.SetChar(' ', cardinal.Position{X: pos.X + x, Y: pos.Y + y})
		}
	}
}

func (c *Console) PrintRectStyle(pos cardinal.Position, w, h uint, boxStyle BorderStyle, filled, clear bool) {
	if clear {
		c.ClearRect(pos, w, h)
	}

	// Background
	if filled {
		for y := 0; y < int(h)-1; y++ {
			for x := 0; x < int(w); x++ {
				c.SetCharBackground(cardinal.Position{X: pos.X + x, Y: pos.Y + y}, c.GetDefaultBackground())
			}
		}
	}

	// Top/Bottom
	for x := uint(pos.X + 1); x < uint(pos.X)+w; x++ {
		c.PutCharEx(boxStyle.H, cardinal.Position{X: int(x), Y: pos.Y}, rl.White, rl.Black)
		c.PutCharEx(boxStyle.H, cardinal.Position{X: int(x), Y: pos.Y + int(h-1)}, rl.White, rl.Black)
	}

	// Left/Right
	for y := uint(0); y < h-1; y++ {
		c.PutCharEx(boxStyle.V, cardinal.Position{X: pos.X, Y: pos.Y + int(y)}, rl.White, rl.Black)
		c.PutCharEx(boxStyle.V, cardinal.Position{X: pos.X + int(w-1), Y: pos.Y + int(y)}, rl.White, rl.Black)
	}

	// Corners
	c.PutCharEx(boxStyle.NE, cardinal.Position{X: pos.X + int(w-1), Y: pos.Y}, rl.White, rl.Black)
	c.PutCharEx(boxStyle.SE, cardinal.Position{X: pos.X + int(w-1), Y: pos.Y + int(h-1)}, rl.White, rl.Black)
	c.PutCharEx(boxStyle.SW, cardinal.Position{X: pos.X, Y: pos.Y + int(h-1)}, rl.White, rl.Black)
	c.PutCharEx(boxStyle.NW, cardinal.Position{X: pos.X, Y: pos.Y}, rl.White, rl.Black)
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
	// @todo
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
