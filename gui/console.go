package gui

import "C"
import (
	"bytes"
	"container/list"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-rogue/engine/geom"
	"github.com/go-rogue/engine/sprites"
	"strings"
)

// FrameStyle defines a Widget's border. Those without a border will have it set to ZeroWallBorder.
type FrameStyle struct {
	V, H, NE, SE, SW, NW uint
}

// SingleWallBorder displays a single line border
var SingleWallBorder = FrameStyle{
	sprites.TCOD_CHAR_VLINE, sprites.TCOD_CHAR_HLINE, sprites.TCOD_CHAR_NE, sprites.TCOD_CHAR_SE, sprites.TCOD_CHAR_SW, sprites.TCOD_CHAR_NW,
}

// ZeroWallBorder displays no border
var ZeroWallBorder = FrameStyle{
	0, 0, 0, 0, 0, 0,
}

// IsZeroWallBorder is used to determine if a Widget's border should be drawn
func (f FrameStyle) IsZeroWallBorder() bool {
	return f.V == 0 && f.H == 0 && f.NE == 0 && f.SE == 0 && f.SW == 0 && f.NW == 0
}

// FrameTitle is complementary to FrameStyle, both are used by IConsole.PrintFrame
type FrameTitle struct {
	text      string
	alignment TextAlignment
}

var ZeroFrameTitle = FrameTitle{
	text:      "",
	alignment: AlignTextCenter,
}

func (f FrameTitle) IsVisible() bool {
	return f.text != ""
}

func (f FrameTitle) Position(homePos geom.Point, width uint) geom.Point {
	if f.alignment == AlignTextLeft {
		return geom.Point{
			X: homePos.X,
			Y: homePos.Y,
		}
	}

	if f.alignment == AlignTextRight {
		return geom.Point{
			X: homePos.X + int(width) - utf8.RuneCountInString(f.text),
			Y: homePos.Y,
		}
	}

	return geom.Point{
		X: homePos.X + int(width/2) - utf8.RuneCountInString(f.text)/2,
		Y: homePos.Y,
	}
}

type CellLayer struct {
	char   uint
	fg, bg rl.Color
}

type Cell struct {
	char   uint
	fg, bg rl.Color
}

type CellMap map[geom.Point]*Cell

type DirtyMap map[geom.Point]bool

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
	GetData() CellMap
	GetDirty() DirtyMap
	ClearDirty()
	GetCellAtPos(pos geom.Point) *Cell
	GetDefaultBackground() rl.Color
	GetDefaultForeground() rl.Color
	SetDefaultForeground(colour rl.Color)
	SetDefaultBackground(colour rl.Color)
	Clear()
	GetCharBackground(pos geom.Point) rl.Color
	GetCharForeground(pos geom.Point) rl.Color
	SetCharBackground(pos geom.Point, colour rl.Color)
	SetCharForeground(pos geom.Point, colour rl.Color)
	SetChar(r uint, pos geom.Point)
	PutChar(r uint, p geom.Point)
	PutCharEx(r uint, p geom.Point, fg, bg rl.Color)
	Print(pos geom.Point, str string)
	//PrintRect(x, y, w, h int, fmts string, v ...interface{}) int
	ClearRect(pos geom.Point, w, h uint)
	PrintFrame(pos geom.Point, w, h uint, boxStyle FrameStyle, filled, clear bool)
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
	GetChar(pos geom.Point) uint
	GetWidth() uint
	GetHeight() uint
	SetKeyColor(colour rl.Color)
	Blit(xSrc, ySrc, wSrc, hSrc int, dst IConsole, xDst, yDst int, foregroundAlpha, backgroundAlpha float32)
}

type Console struct {
	width, height        uint
	data                 CellMap
	dirty                DirtyMap
	defaultBg, defaultFg rl.Color
}

// Initiates the CellMap
func (c *Console) init() {
	c.data = make(CellMap)
	c.dirty = make(map[geom.Point]bool)

	for cY := 0; cY < int(c.height); cY++ {
		for cX := 0; cX < int(c.width); cX++ {
			c.PutChar(' ', geom.Point{X: cX, Y: cY})
		}
	}
}

func (c Console) GetDirty() DirtyMap {
	return c.dirty
}

func (c *Console) SetDirty(pos geom.Point) {
	c.dirty[pos] = true
}

func (c *Console) ClearDirty() {
	for k := range c.dirty {
		delete(c.dirty, k)
	}
}

func (c Console) GetData() CellMap {
	return c.data
}

func (c *Console) GetCellAtPos(pos geom.Point) *Cell {
	return c.data[pos]
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
	c.ClearDirty()

	for k := range c.data {
		delete(c.data, k)
	}

	for cY := 0; cY < int(c.height); cY++ {
		for cX := 0; cX < int(c.width); cX++ {
			c.PutChar(' ', geom.Point{X: cX, Y: cY})
		}
	}
}

// TODO This will result in a panic if pos is out of bounds
func (c *Console) SetChar(r uint, pos geom.Point) {
	c.data[pos].char = r
	c.SetDirty(pos)
}

// TODO This will result in a panic if pos is out of bounds
func (c Console) GetCharBackground(pos geom.Point) rl.Color {
	return c.data[pos].bg
}

// TODO This will result in a panic if pos is out of bounds
func (c Console) GetCharForeground(pos geom.Point) rl.Color {
	return c.data[pos].fg
}

// TODO This will result in a panic if pos is out of bounds
func (c *Console) SetCharBackground(pos geom.Point, colour rl.Color) {
	c.data[pos].bg = colour
	c.SetDirty(pos)
}

// TODO This will result in a panic if pos is out of bounds
func (c *Console) SetCharForeground(pos geom.Point, colour rl.Color) {
	c.data[pos].fg = colour
	c.SetDirty(pos)
}

// TODO This will result in a panic if p is out of bounds
func (c *Console) PutChar(r uint, p geom.Point) {
	c.data[p] = &Cell{char: r, fg: c.defaultFg, bg: sprites.ColourNC}
	c.SetDirty(p)
}

// TODO This will result in a panic if p is out of bounds
func (c *Console) PutCharEx(r uint, p geom.Point, fg, bg rl.Color) {
	c.data[p] = &Cell{char: r, fg: fg, bg: bg}
	c.SetDirty(p)
}

//
// Split str into individual runes and then loop over and set the
// char at that geom position.
//
// <%FG:colour_name>Text<%/>
// <%BG:colour_name>Text<%/>
// <%FG:colour_name,BG:colour_name>Text<%/>
//
func (c *Console) Print(pos geom.Point, str string) {
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
			c.PutCharEx(uint(r), geom.Point{X: pos.X + i + xOff, Y: pos.Y}, fg, bg)
		}
	}
}

//
// Reset cells within the Rectangular area to default
//
func (c *Console) ClearRect(pos geom.Point, w, h uint) {
	for y := 0; y < int(h)-1; y++ {
		for x := 0; x < int(w); x++ {
			c.SetChar(' ', geom.Point{X: pos.X + x, Y: pos.Y + y})
		}
	}
}

func (c *Console) PrintFrame(pos geom.Point, w, h uint, style FrameStyle, filled, clear bool) {
	if clear {
		c.ClearRect(pos, w, h)
	}

	// Background
	if filled {
		for y := 0; y < int(h)-1; y++ {
			for x := 0; x < int(w); x++ {
				c.SetCharBackground(geom.Point{X: pos.X + x, Y: pos.Y + y}, c.GetDefaultBackground())
			}
		}
	}

	// Top/Bottom
	for x := uint(pos.X + 1); x < uint(pos.X)+w; x++ {
		c.PutCharEx(style.H, geom.Point{X: int(x), Y: pos.Y}, rl.White, rl.Black)
		c.PutCharEx(style.H, geom.Point{X: int(x), Y: pos.Y + int(h-1)}, rl.White, rl.Black)
	}

	// Left/Right
	for y := uint(0); y < h-1; y++ {
		c.PutCharEx(style.V, geom.Point{X: pos.X, Y: pos.Y + int(y)}, rl.White, rl.Black)
		c.PutCharEx(style.V, geom.Point{X: pos.X + int(w-1), Y: pos.Y + int(y)}, rl.White, rl.Black)
	}

	// Corners
	c.PutCharEx(style.NE, geom.Point{X: pos.X + int(w-1), Y: pos.Y}, rl.White, rl.Black)
	c.PutCharEx(style.SE, geom.Point{X: pos.X + int(w-1), Y: pos.Y + int(h-1)}, rl.White, rl.Black)
	c.PutCharEx(style.SW, geom.Point{X: pos.X, Y: pos.Y + int(h-1)}, rl.White, rl.Black)
	c.PutCharEx(style.NW, geom.Point{X: pos.X, Y: pos.Y}, rl.White, rl.Black)
}

func (c Console) GetChar(pos geom.Point) uint {
	return c.data[pos].char
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
