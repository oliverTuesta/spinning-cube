package math

import (
	"math"

	"github.com/gdamore/tcell/v2"
)

func DrawLine3D(x1, y1, z1, x2, y2, z2 float64, screen tcell.Screen) {
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

