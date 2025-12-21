package math3d

import "math"

type Cube struct {
	Vertices []Vec3
	Edges    [][2]int
}

func NewCube(width float64) Cube {
	return Cube{
		Vertices: []Vec3{
			{-width, -width, -width},
			{width, -width, -width},
			{width, width, -width},
			{-width, width, -width},
			{-width, -width, width},
			{width, -width, width},
			{width, width, width},
			{-width, width, width},
		},
		Edges: [][2]int{
			{0, 1}, {1, 2}, {2, 3}, {3, 0},
			{4, 5}, {5, 6}, {6, 7}, {7, 4},
			{0, 4}, {1, 5}, {2, 6}, {3, 7},
		},
	}
}

func (c *Cube) Move(x float64, y float64, z float64) {
	for i := 0; i < len(c.Vertices); i++ {
		c.Vertices[i].x += x
		c.Vertices[i].y += y
		c.Vertices[i].z += z
	}
}

func (c *Cube) multiplyMatrix(matrix [][]float64) {
	for i, vec := range c.Vertices {
		var vals []float64
		for j := 0; j < len(matrix); j++ {
			v := matrix[j][0] * vec.x + matrix[j][1] * vec.y + matrix[j][2] *	vec.z
			vals = append(vals, v)
		}
		c.Vertices[i] = Vec3 {x: vals[0], y: vals[1], z: vals[2]}
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
		{math.Cos(rad), math.Sin(rad), 0},
		{math.Sin(rad), math.Cos(rad), 0},
		{0, 0, 1},
	}
	c.multiplyMatrix(matrix)
}
