package sprites

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileSetProperties struct {
	Filename string
	Codec    Codec
	W, H     uint
}

//noinspection GoSnakeCaseUsage,GoUnusedConst
const (
	/* single walls */
	TCOD_CHAR_HLINE = 196
	TCOD_CHAR_VLINE = 179
	TCOD_CHAR_NE    = 191
	TCOD_CHAR_NW    = 218
	TCOD_CHAR_SE    = 217
	TCOD_CHAR_SW    = 192
	TCOD_CHAR_TEEW  = 180
	TCOD_CHAR_TEEE  = 195
	TCOD_CHAR_TEEN  = 193
	TCOD_CHAR_TEES  = 194
	TCOD_CHAR_CROSS = 197
	/* double walls */
	TCOD_CHAR_DHLINE = 205
	TCOD_CHAR_DVLINE = 186
	TCOD_CHAR_DNE    = 187
	TCOD_CHAR_DNW    = 201
	TCOD_CHAR_DSE    = 188
	TCOD_CHAR_DSW    = 200
	TCOD_CHAR_DTEEW  = 185
	TCOD_CHAR_DTEEE  = 204
	TCOD_CHAR_DTEEN  = 202
	TCOD_CHAR_DTEES  = 203
	TCOD_CHAR_DCROSS = 206
	/* blocks */
	TCOD_CHAR_BLOCK1 = 176
	TCOD_CHAR_BLOCK2 = 177
	TCOD_CHAR_BLOCK3 = 178
	/* arrows */
	TCOD_CHAR_ARROW_N = 24
	TCOD_CHAR_ARROW_S = 25
	TCOD_CHAR_ARROW_E = 26
	TCOD_CHAR_ARROW_W = 27
	/* arrows without tail */
	TCOD_CHAR_ARROW2_N = 30
	TCOD_CHAR_ARROW2_S = 31
	TCOD_CHAR_ARROW2_E = 16
	TCOD_CHAR_ARROW2_W = 17
	/* double arrows */
	TCOD_CHAR_DARROW_H = 29
	TCOD_CHAR_DARROW_V = 18
	/* GUI stuff */
	TCOD_CHAR_CHECKBOX_UNSET = 224
	TCOD_CHAR_CHECKBOX_SET   = 225
	TCOD_CHAR_RADIO_UNSET    = 9
	TCOD_CHAR_RADIO_SET      = 10
	/* sub-pixel resolution kit */
	TCOD_CHAR_SUBP_NW   = 226
	TCOD_CHAR_SUBP_NE   = 227
	TCOD_CHAR_SUBP_N    = 228
	TCOD_CHAR_SUBP_SE   = 229
	TCOD_CHAR_SUBP_DIAG = 230
	TCOD_CHAR_SUBP_E    = 231
	TCOD_CHAR_SUBP_SW   = 232
	/* miscellaneous */
	TCOD_CHAR_SMILIE         = 1
	TCOD_CHAR_SMILIE_INV     = 2
	TCOD_CHAR_HEART          = 3
	TCOD_CHAR_DIAMOND        = 4
	TCOD_CHAR_CLUB           = 5
	TCOD_CHAR_SPADE          = 6
	TCOD_CHAR_BULLET         = 7
	TCOD_CHAR_BULLET_INV     = 8
	TCOD_CHAR_MALE           = 11
	TCOD_CHAR_FEMALE         = 12
	TCOD_CHAR_NOTE           = 13
	TCOD_CHAR_NOTE_DOUBLE    = 14
	TCOD_CHAR_LIGHT          = 15
	TCOD_CHAR_EXCLAM_DOUBLE  = 19
	TCOD_CHAR_PILCROW        = 20
	TCOD_CHAR_SECTION        = 21
	TCOD_CHAR_POUND          = 156
	TCOD_CHAR_MULTIPLICATION = 158
	TCOD_CHAR_FUNCTION       = 159
	TCOD_CHAR_RESERVED       = 169
	TCOD_CHAR_HALF           = 171
	TCOD_CHAR_ONE_QUARTER    = 172
	TCOD_CHAR_COPYRIGHT      = 184
	TCOD_CHAR_CENT           = 189
	TCOD_CHAR_YEN            = 190
	TCOD_CHAR_CURRENCY       = 207
	TCOD_CHAR_THREE_QUARTERS = 243
	TCOD_CHAR_DIVISION       = 246
	TCOD_CHAR_GRADE          = 248
	TCOD_CHAR_UMLAUT         = 249
	TCOD_CHAR_POW1           = 251
	TCOD_CHAR_POW3           = 252
	TCOD_CHAR_POW2           = 253
	TCOD_CHAR_BULLET_SQUARE  = 254
	/* diacritics */
)

//
// The tcod_codec_ comes from https://github.com/libtcod/libtcod/blob/master/src/libtcod/sys_sdl_c.cpp#L165
// It is the codec for TCOD_FONT_LAYOUT_TCOD and converts from EASCII code-point -> raw tile Position.
// BSD 3-Clause License
// Copyright © 2008-2019, Jice and the libtcod contributors. All rights reserved.
//

var TCODCodec = [256]uint{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 76, 77, 0, 0, 0, 0, 0, /* 0 to 15 */
	71, 70, 72, 0, 0, 0, 0, 0, 64, 65, 67, 66, 0, 73, 68, 69, /* 16 to 31 */
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, /* 32 to 47 */
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, /* 48 to 63 */
	32, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, /* 64 to 79 */
	111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 33, 34, 35, 36, 37, /* 80 to 95 */
	38, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, /* 96 to 111 */
	143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 39, 40, 41, 42, 0, /* 112 to 127 */
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, /* 128 to 143 */
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, /* 144 to 159 */
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, /* 160 to 175 */
	43, 44, 45, 46, 49, 0, 0, 0, 0, 81, 78, 87, 88, 0, 0, 55, /* 176 to 191 */
	53, 50, 52, 51, 47, 48, 0, 0, 85, 86, 82, 84, 83, 79, 80, 0, /* 192 to 207 */
	0, 0, 0, 0, 0, 0, 0, 0, 0, 56, 54, 0, 0, 0, 0, 0, /* 208 to 223 */
	74, 75, 57, 58, 59, 60, 61, 62, 63, 0, 0, 0, 0, 0, 0, 0, /* 224 to 239 */
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, /* 240 to 255 */
}

type Codec int

const (
	LayoutNone Codec = iota
	LayoutTcod
	LayoutAsciiInCol
	LayoutAsciiInRow
)

type Tileset struct {
	sprites  SpriteSheet
	asciiMap map[uint]uint // ascii code to idx of sprite in sheet
	max      uint          // Max codepoint value
}

func (t *Tileset) decode(c Codec) {
	if c == LayoutNone {
		t.max = t.sprites.MaxIdx()

		for i := uint(0); i < t.max; i++ {

		}
	} else if c == LayoutAsciiInCol || c == LayoutAsciiInRow {
		panic("Not implemented")
	} else if c == LayoutTcod {
		t.max = 256
		for i := uint(0); i < t.max; i++ {
			t.mapAsciiToFont(i, TCODCodec[i], 0)
		}
		t.mapClone(0x2500, TCOD_CHAR_HLINE)
		t.mapClone(0x2502, TCOD_CHAR_VLINE)
		t.mapClone(0x250C, TCOD_CHAR_NW)
		t.mapClone(0x2510, TCOD_CHAR_NE)
		t.mapClone(0x2514, TCOD_CHAR_SW)
		t.mapClone(0x2518, TCOD_CHAR_SE)
	}
}

func (t *Tileset) mapAsciiToFont(asciiCode, fontCharX, fontCharY uint) {
	t.asciiMap[asciiCode] = t.sprites.IdxAt(fontCharX, fontCharY)
}

func (t *Tileset) mapClone(newCodepoint, oldCodepoint uint) {
	if oldCodepoint >= t.max {
		return
	}

	t.mapAsciiToFont(newCodepoint, t.asciiMap[oldCodepoint], 0)
}

func (t Tileset) Debug() {
	for i := uint(0); i < t.max; i++ {
		fmt.Println("ASCII [", i, "] to idx [", t.asciiMap[i], "]")
	}
}

func (t Tileset) GetSpriteForChar(c uint) *Sprite {
	if c > 256 {
		c = 0
	}
	return t.sprites.AtIdx(t.asciiMap[c])
}

func (t *Tileset) Unload() {
	t.sprites.Unload()
}

func (t Tileset) GetTileWidth() uint {
	return t.sprites.TileWidth
}

func (t Tileset) GetTileHeight() uint {
	return t.sprites.TileHeight
}

func (t Tileset) GetTexture2D() rl.Texture2D {
	return t.sprites.TxTiles
}

//
// @todo should w/h be inferred here from the filename?
// @todo if filename doesnt exist return an error!
//
func NewTileSet(filename string, c Codec, w, h uint) *Tileset {
	tex := rl.LoadTexture(filename)
	tileset := &Tileset{
		sprites:  *NewSpriteSheet(tex, w, h),
		asciiMap: make(map[uint]uint),
	}

	tileset.decode(c)

	return tileset
}

func NewTileSetFromProps(conf TileSetProperties) *Tileset {
	return NewTileSet(conf.Filename, conf.Codec, conf.W, conf.H)
}
