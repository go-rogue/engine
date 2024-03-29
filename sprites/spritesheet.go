package sprites

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-rogue/engine/geom"
	"math"
)

type Sprite struct {
	R rl.Rectangle
	t *SpriteSheet
}

func (s Sprite) Draw(pos geom.Point, fg, bg rl.Color) {
	if bg != ColourNC {
		tileWidth := int(s.t.TileWidth)
		tileHeight := int(s.t.TileHeight)
		// @todo this draw call burns the cpu
		// I think its because it gets called once for each tile on the grid
		rl.DrawRectangle(int32(pos.X*tileWidth), int32(pos.Y*tileHeight), int32(tileWidth), int32(tileHeight), bg)
	}

	rl.DrawTextureRec(s.t.TxTiles, s.R, pos.Vector2(int(s.t.TileWidth), int(s.t.TileHeight)), fg)
}

type SpriteSheet struct {
	TxTiles    rl.Texture2D // Sprite sheet texture
	TileWidth  uint
	TileHeight uint
	Cols       uint
	Rows       uint
	Tiles      []*Sprite
}

func NewSpriteSheet(tx rl.Texture2D, w, h uint) *SpriteSheet {
	cols := uint(math.Ceil(float64(tx.Width / int32(w))))
	rows := uint(math.Ceil(float64(tx.Height / int32(h))))

	tileSheet := &SpriteSheet{
		TxTiles:    tx,
		TileWidth:  w,
		TileHeight: h,
		Cols:       cols,
		Rows:       rows,
		Tiles:      make([]*Sprite, cols*rows),
	}

	// 320 x 90
	for y := uint(0); y < rows; y++ {
		for x := uint(0); x < cols; x++ {
			tileSheet.Set(x, y, &Sprite{
				R: rl.NewRectangle(float32(x*w), float32(y*h), float32(w), float32(h)),
				t: tileSheet,
			})
			// fmt.Println("Setting Sprite (", x, ",", y, ") Idx (", tileSheet.IdxAt(x, y), ") as Rect(", x*w, ",", y*h, ",", (x*w)+10, ",", (y*h)+10, ")")
		}
	}

	return tileSheet
}

func (t *SpriteSheet) Unload() {
	rl.UnloadTexture(t.TxTiles)
}

func (t SpriteSheet) At(x, y uint) *Sprite {
	return t.Tiles[t.IdxAt(x, y)]
}

func (t SpriteSheet) AtIdx(idx uint) *Sprite {
	// @todo add bounds check
	return t.Tiles[idx]
}

func (t SpriteSheet) IdxAt(x, y uint) uint {
	return (t.Cols * y) + x
}

func (t *SpriteSheet) Set(x, y uint, val *Sprite) {
	t.Tiles[t.IdxAt(x, y)] = val
}

func (t SpriteSheet) MaxIdx() uint {
	return t.Cols * t.Rows
}
