package math2d

type Vec2 struct {
	X, Y float64
}


func (a Vec2) Sub(b Vec2) Vec2 {
	return Vec2 {
		a.X - b.X,
		a.Y - b.Y,
	}
}

func (a Vec2) Cross(b Vec2) float64 {
    return a.X*b.Y - a.Y*b.X
}
