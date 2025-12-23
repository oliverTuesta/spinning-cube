package math2d

type Square struct {
	vertices []Vec2
	edges [][2]int
}

func NewSquare(width float64) Square {
	return Square{
		vertices: []Vec2 {
			{0, 0},
			{width, 0},
			{width, width},
			{0, width},
		},
		edges: [][2]int {
			{0,1},
			{1,2},
			{2,3},
			{3,0},
		},
	}
}

func (s *Square) Move(x float64, y float64) {
	for i := 0; i < len(s.vertices); i++ {
		s.vertices[i].X += x
		s.vertices[i].Y += y
	}
}

func (s *Square) multiplyMatrix(matrix [][2]float64) {
	for i, vec := range s.vertices {
		var vals []float64
		for j := 0; j < len(matrix); j++ {
			v := matrix[j][0] * vec.X + matrix[j][1] * vec.Y
			vals = append(vals, v)
		}
		s.vertices[i] = Vec2 {X: vals[0], Y: vals[1]}
	}
}

