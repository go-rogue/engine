package geom

import rl "github.com/gen2brain/raylib-go/raylib"

// Point is a data structure for storing integer geom coordinates.
type Point struct {
	X, Y int
}

// Vector2 provides the dot product of Point, used for getting Rect size for texture display.
func (p Point) Vector2(expX, expY int) rl.Vector2 {
	return rl.Vector2{X: float32(p.X * expX), Y: float32(p.Y * expY)}
}

// DiagonalNeighbours generates all Point that share a corner with self. (NE,SE,SW,NW)
func (p Point) DiagonalNeighbours() []Point {
	return nil // TODO...
}

// Neighbours generates all Point that share a horizontal or vertical edge with this one. (N,S,E,W)
func (p Point) Neighbours() []Point {
	return nil // TODO...
}

// GetClosestPosition of all the positions in candidates, return the one closest to self.
func (p Point) GetClosestPosition(candidates []Point) Point {
	return Point{} // TODO...
}

// GetFarthestPosition of all the positions in candidates, return the one farthest from self.
func (p Point) GetFarthestPosition(candidates []Point) Point {
	return Point{} // TODO...
}

// ManhattanDistanceTo returns Manhattan distance between self and target Point.
func (p Point) ManhattanDistanceTo(point Point) int {
	return 0 // TODO...
}

// PathLTo generates all positions on an L-shaped path between self and target Point.
func (p Point) PathLTo(point Point) []Point {
	return nil // TODO...
}

// PositionsBresenhamTo generates all positions on a Bresenham algorithm line between self and other Point.
func (p Point) PositionsBresenhamTo(point Point) []Point {
	return nil // TODO...
}
