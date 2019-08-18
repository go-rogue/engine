package geom

import rl "github.com/gen2brain/raylib-go/raylib"

//
// Data structure for storing geom coordinates
//
type Position struct {
	X, Y int
}

func (pos Position) Vector2(expX, expY int) rl.Vector2 {
	return rl.Vector2{X: float32(pos.X * expX), Y: float32(pos.Y * expY)}
}

//
// Generates all points that share a corner with self (NE,SE,SW,NW)
//
func (pos Position) DiagonalNeighbours() []Position {
	return nil
}

//
// Generates all points that share a horizontal or vertical edge with this one. (N,S,E,W)
//
func (pos Position) Neighbours() []Position {
	return nil
}

//
// Of all the positions in candidates, return the one closest to self.
//
func (pos Position) GetClosestPosition(candidates []Position) Position {
	return Position{}
}

//
// Of all the positions in candidates, return the one farthest from self.
//

func (pos Position) GetFarthestPosition(candidates []Position) Position {
	return Position{}
}

//
// Returns Manhattan distance between self and target
//

func (pos Position) ManhattanDistanceTo(p Position) int {
	return 0
}

//
// Generates all positions on an L-shaped path between self and target.
//

func (pos Position) PathLTo(p Position) []Position {
	return nil
}

//
// Generates all positions on a Bresenham algorithm line between self and other.
//

func (pos Position) PositionsBresenhamTo(p Position) []Position {
	return nil
}
