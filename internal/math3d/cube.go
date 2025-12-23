package math3d

import "math"

type Cube struct {
	vertices []Vec3
	edges    [][2]int
}

func NewCube(width float64) Cube {
	return Cube{
		vertices: []Vec3{
			{0, 0, 0},
			{width, 0, 0},
			{width, width, 0},
			{0, width, 0},
			{0, 0, width},
			{width, 0, width},
			{width, width, width},
			{0, width, width},
		},
		edges: [][2]int{
			{0, 1}, {1, 2}, {2, 3}, {3, 0},
			{4, 5}, {5, 6}, {6, 7}, {7, 4},
			{0, 4}, {1, 5}, {2, 6}, {3, 7},
		},
	}
}

func (c *Cube) Move(x float64, y float64, z float64) {
	for i := 0; i < len(c.vertices); i++ {
		c.vertices[i].x += x
		c.vertices[i].y += y
		c.vertices[i].z += z
	}
}

func (c *Cube) multiplyMatrix(matrix [][]float64) {
	for i, vec := range c.vertices {
		var vals []float64
		for j := 0; j < len(matrix); j++ {
			v := matrix[j][0] * vec.x + matrix[j][1] * vec.y + matrix[j][2] *	vec.z
			vals = append(vals, v)
		}
		c.vertices[i] = Vec3 {x: vals[0], y: vals[1], z: vals[2]}
	}
}

func (c *Cube) RotateX(rad float64) {
	matrix :=  [][]float64 {
		{1, 0, 0},
		{0, math.Cos(rad), -math.Sin(rad)},
		{0, math.Sin(rad), math.Cos(rad)},
	}
	c.multiplyMatrix(matrix)
}

func (c *Cube) RotateY(rad float64) {
	matrix :=  [][]float64 {
		{math.Cos(rad), 0, math.Sin(rad)},
		{0, 1, 0},
		{-math.Sin(rad), 0, math.Cos(rad)},
	}
	c.multiplyMatrix(matrix)
}

func (c *Cube) RotateZ(rad float64) {
	matrix :=  [][]float64 {
		{math.Cos(rad), -math.Sin(rad), 0},
		{math.Sin(rad), math.Cos(rad), 0},
		{0, 0, 1},
	}
	c.multiplyMatrix(matrix)
}

func (c *Cube) Center() Vec3 {
    var sum Vec3
    for _, v := range c.vertices {
        sum.x += v.x
        sum.y += v.y
        sum.z += v.z
    }
    n := float64(len(c.vertices))
    return Vec3{sum.x / n, sum.y / n, sum.z / n}
}

func (c *Cube) RotateAroundCenter(rx, ry, rz float64) {
    center := c.Center()
    c.Move(-center.x, -center.y, -center.z)
    c.RotateY(ry)
    c.RotateX(rx)
    c.RotateZ(rz)
    c.Move(center.x, center.y, center.z)
}
