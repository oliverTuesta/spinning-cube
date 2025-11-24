package math2d

import (
	"math"

	"github.com/gdamore/tcell/v2"
)

func Rotate(x float64, y float64, t float64) (float64, float64) {
	matrix := [][]float64{
		{math.Cos(t), -math.Sin(t)},
		{math.Sin(t), math.Cos(t)},
	}

	vec := []float64{x, y}

	result := make([]float64, 2)

	for i := range len(matrix) {
		for j := range len(vec) {
			result[i] += matrix[i][j] * vec[j]
		}
	}

	return result[0], result[1]
}

func DrawLine(x1, y1, x2, y2 float64, screen tcell.Screen) {
	steps := int(math.Max(math.Abs(x2-x1), math.Abs(y2-y1)))

	dx := (x2 - x1) / float64(steps)
	dy := (y2 - y1) / float64(steps)

	x := x1
	y := y1

	style := tcell.StyleDefault.Foreground(tcell.ColorGreen)

	for i := 0; i <= steps; i++ {
		screen.SetContent(int(math.Round(x)), int(math.Round(y)), '*', nil, style)
		x += dx
		y += dy
	}
}

func DrawRectangle(coords [4][2]float64, screen tcell.Screen, cx float64, cy float64) {
	for i := range len(coords) - 1 {
		DrawLine(coords[i][0]+cx, (coords[i][1]/2.0)+cy, coords[i+1][0]+cx, (coords[i+1][1]/2.0)+cy, screen)
	}
	DrawLine(coords[3][0]+cx, (coords[3][1]/2.0)+cy, coords[0][0]+cx, (coords[0][1]/2.0)+cy, screen)
}
