package math3d

import (
	"github.com/gdamore/tcell/v2"
	"github.com/olivertuesta/spinning-cube/internal/math2d"
)


type Graphics3D struct {
	Screen *tcell.Screen
	CameraDist float64
	G2 *math2d.Graphics2D 
}

func NewGraphics3D(screen *tcell.Screen, cd float64) (Graphics3D) {
	return Graphics3D {Screen: screen, CameraDist: cd, G2: math2d.NewGraphics2D(screen) }
}

func (g *Graphics3D) DrawVectors(v1 Vec3, v2 Vec3) {
	g.G2.DrawLine(v1.toVec2(g.CameraDist), v2.toVec2(g.CameraDist))
}

func (g *Graphics3D) DrawCube(c *Cube) {
	for _, edge := range c.Edges {
		g.DrawVectors(c.Vertices[edge[0]], c.Vertices[edge[1]])
	}
}
