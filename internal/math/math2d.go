package math

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

func DrawLine2D(x1, y1, x2, y2 float64, screen tcell.Screen) {
	// m := (y2 - y1) / (x2 - x1)
	// b := y1 - m*x1
	steps := math.Round(max(math.Abs(x1-x2), math.Abs(y1-y2)))

	style := tcell.StyleDefault.Foreground(tcell.ColorGreen)

	x := x1
	y := y1

	for i := 0; i < int(steps); i++ {
		screen.SetContent(int(math.Round(x)), int(math.Round(y)), '*', nil, style)
		x += (x2 - x1) / steps
		y += (y2 - y1) / steps
	}
}

func DrawRectangle(coords [4][2]float64, screen tcell.Screen, cx float64, cy float64) {
	for i := range len(coords) - 1 {
		DrawLine2D(coords[i][0]+cx, (coords[i][1]/2.0)+cy, coords[i+1][0]+cx, (coords[i+1][1]/2.0)+cy, screen)
	}
	DrawLine2D(coords[3][0]+cx, (coords[3][1]/2.0)+cy, coords[0][0]+cx, (coords[0][1]/2.0)+cy, screen)
}
