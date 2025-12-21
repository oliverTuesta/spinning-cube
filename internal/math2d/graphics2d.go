package math2d

import (
	"math"

	"github.com/gdamore/tcell/v2"
)

type Graphics2D struct {
	Screen *tcell.Screen
}

func NewGraphics2D (screen *tcell.Screen) (*Graphics2D){
	return &Graphics2D { Screen: screen }
}

func (g *Graphics2D) DrawLine(v1 Vec2, v2 Vec2) {
	steps := math.Max(1, math.Round(
		math.Max(math.Abs(v1.X-v2.X), math.Abs(v1.Y-v2.Y)),
	))

	style := tcell.StyleDefault.Foreground(tcell.ColorGreen)

	w, h := (*g.Screen).Size()
	cx := float64(w) / 2
	cy := float64(h) / 2

	v := v1

	for i := 0; i < int(steps); i++ {
		sx := int(math.Round(v.X + cx))
		sy := int(math.Round(v.Y/2 + cy))

		(*g.Screen).SetContent(sx, sy, '*', nil, style)

		v.X += (v2.X - v1.X) / steps
		v.Y += (v2.Y - v1.Y) / steps
	}
}

