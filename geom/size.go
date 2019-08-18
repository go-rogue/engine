package geom

type Size struct {
	Width, Height int
}

func (s Size) Area() int {
	return s.Width * s.Height
}
