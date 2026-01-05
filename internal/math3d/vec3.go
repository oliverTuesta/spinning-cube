package math3d

import (
	"math"

	"github.com/olivertuesta/spinning-cube/internal/math2d"
)

type Vec3 struct {
	X,Y,Z float64
}

func (a Vec3) toVec2(cameraDist float64) math2d.Vec2 {
	ad := cameraDist + a.Z

	op := a.Y
	rad := math.Atan(op / ad)
	y := math.Tan(rad) * cameraDist

	op = a.X
	rad = math.Atan(op / ad)
	x := math.Tan(rad) * cameraDist

	return math2d.Vec2{X: x, Y: y}
}

func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
	}
}

func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

func (a Vec3) Dot(b Vec3) float64 {
	return a.X * b.X + a.Y * b.Y + a.Z * b.Z
}

func (v Vec3) Normalize() Vec3 {
	len := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	return Vec3{
		v.X / len,
		v.Y / len,
		v.Z / len,
	}
}
