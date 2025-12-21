package math3d

import (
	"math"

	"github.com/olivertuesta/spinning-cube/internal/math2d"
)

type Vec3 struct {
	x, y, z float64
}

func (v Vec3) toVec2(cameraDist float64) math2d.Vec2 {
	ad := cameraDist + v.z

	op := v.y
	rad := math.Atan(op / ad)
	y := math.Tan(rad) * cameraDist

	op = v.x
	rad = math.Atan(op / ad)
	x := math.Tan(rad) * cameraDist

	return math2d.Vec2{X: x, Y: y}
}
