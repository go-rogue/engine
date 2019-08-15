package cardinal

import rl "github.com/gen2brain/raylib-go/raylib"

//
// Data structure for storing cardinal coordinates
//
type Position struct {
	X, Y int
}

func (pos Position) Vector2(expX, expY int) rl.Vector2 {
	return rl.Vector2{X: float32(pos.X * expX), Y: float32(pos.Y * expY)}
}
