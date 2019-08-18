package geom

import rl "github.com/gen2brain/raylib-go/raylib"

//
// Data structure for storing geom coordinates
//
type Point struct {
	X, Y int
}

func (p Point) Vector2(expX, expY int) rl.Vector2 {
	return rl.Vector2{X: float32(p.X * expX), Y: float32(p.Y * expY)}
}

//
// Generates all points that share a corner with self (NE,SE,SW,NW)
//
func (p Point) DiagonalNeighbours() []Point {
	return nil
}

//
// Generates all points that share a horizontal or vertical edge with this one. (N,S,E,W)
//
func (p Point) Neighbours() []Point {
	return nil
}

//
// Of all the positions in candidates, return the one closest to self.
//
func (p Point) GetClosestPosition(candidates []Point) Point {
	return Point{}
}

//
// Of all the positions in candidates, return the one farthest from self.
//

func (p Point) GetFarthestPosition(candidates []Point) Point {
	return Point{}
}

//
// Returns Manhattan distance between self and target
//

func (p Point) ManhattanDistanceTo(point Point) int {
	return 0
}

//
// Generates all positions on an L-shaped path between self and target.
//

func (p Point) PathLTo(point Point) []Point {
	return nil
}

//
// Generates all positions on a Bresenham algorithm line between self and other.
//

func (p Point) PositionsBresenhamTo(point Point) []Point {
	return nil
}
