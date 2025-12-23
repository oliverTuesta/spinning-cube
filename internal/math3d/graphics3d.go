package math3d

import (
	"github.com/gdamore/tcell/v2"
	"github.com/olivertuesta/spinning-cube/internal/math2d"
)


type Graphics3D struct {
	screen *tcell.Screen
	cameraDist float64
	g2 *math2d.Graphics2D 
}

func NewGraphics3D(screen *tcell.Screen, cd float64) (Graphics3D) {
	return Graphics3D {screen: screen, cameraDist: cd, g2: math2d.NewGraphics2D(screen) }
}

func (g *Graphics3D) DrawVectors(v1 Vec3, v2 Vec3) {
	g.g2.DrawLine(v1.toVec2(g.cameraDist), v2.toVec2(g.cameraDist))
}

func (g *Graphics3D) DrawCube(c *Cube) {
	for _, edge := range c.edges {
		g.DrawVectors(c.vertices[edge[0]], c.vertices[edge[1]])
	}
}
