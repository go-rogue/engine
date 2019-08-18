package geom

type Rect struct {
	origin         Point
	size           Size
	x1, x2, y1, y2 int
}

func NewRect(o Point, s Size) *Rect {
	return &Rect{
		origin: o,
		size:   s,
		x1:     o.X,
		x2:     o.X + s.Width,
		y1:     o.Y,
		y2:     o.Y + s.Height,
	}
}

func (r Rect) Origin() Point {
	return r.origin
}

func (r Rect) Size() Size {
	return r.size
}

func (r Rect) Area() int {
	return r.size.Area()
}

//
// Returns the center point of this rectangle
//
func (r Rect) Center() (int, int) {
	return int((r.x1 + r.x2) / 2), int((r.y1 + r.y2) / 2)
}

//
// Returns all points found within this rectangle
//
func (r Rect) Points() []Point {
	points := make([]Point, r.Area())
	i := 0
	for y := r.y1; y < r.y2; y++ {
		for x := r.x1; x < r.x2; x++ {
			points[i] = Point{X: x, Y: y}
		}
	}
	return points
}

//
// Returns true if this rectangle contains all the points found in other
//
func (r Rect) Contains(other Rect) bool {
	return false
}

//
// Returns true if this rectangle intersects with another one
//
func (r Rect) Intersect(other Rect) bool {
	return r.x1 <= other.x2 && r.x2 >= other.x1 && r.y1 <= other.y2 && r.y2 >= other.y1
}
