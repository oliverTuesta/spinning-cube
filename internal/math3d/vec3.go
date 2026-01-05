package math3d

import (
	"math"

	"github.com/olivertuesta/spinning-cube/internal/math2d"
)

type Vec3 struct {
	x,y,z float64
}

func (a Vec3) toVec2(cameraDist float64) math2d.Vec2 {
	ad := cameraDist + a.z

	op := a.y
	rad := math.Atan(op / ad)
	y := math.Tan(rad) * cameraDist

	op = a.x
	rad = math.Atan(op / ad)
	x := math.Tan(rad) * cameraDist

	return math2d.Vec2{X: x, Y: y}
}

func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{
		a.x - b.x,
		a.y - b.y,
		a.z - b.z,
	}
}

func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{
		a.y*b.z - a.z*b.y,
		a.z*b.x - a.x*b.z,
		a.x*b.y - a.y*b.x,
	}
}

func (a Vec3) Dot(b Vec3) float64 {
	return a.x * b.x + a.y * b.y + a.z * b.z
}

func (v Vec3) Normalize() Vec3 {
	len := math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
	return Vec3{
		v.x / len,
		v.y / len,
		v.z / len,
	}
}
