package geom

// Size is a data structure for storing width and height values.
type Size struct {
	Width, Height int
}

// Area returns the calculates area for this Size.
func (s Size) Area() int {
	return s.Width * s.Height
}
