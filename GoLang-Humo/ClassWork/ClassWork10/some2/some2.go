package some2

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) scale(factor float64) {
	r.Height *= factor
	r.Width *= factor
}
