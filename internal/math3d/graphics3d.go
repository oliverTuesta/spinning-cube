package math3d

import (
	"math"

	"github.com/gdamore/tcell/v2"
	"github.com/olivertuesta/spinning-cube/internal/math2d"
)


type Graphics3D struct {
	screen *tcell.Screen
	cameraDist float64
	g2 *math2d.Graphics2D 
	light Vec3
	zbuffer *ZBuffer
}

func NewGraphics3D(screen *tcell.Screen, cd float64, zbuffer *ZBuffer, light Vec3) (Graphics3D) {
	return Graphics3D {
		screen: screen, 
		cameraDist: cd, 
		g2: math2d.NewGraphics2D(screen),
		light: light,
		zbuffer: zbuffer,
	}
}

func (g *Graphics3D) DrawEdge(v1 Vec3, v2 Vec3) {
	g.g2.DrawLine(v1.toVec2(g.cameraDist), v2.toVec2(g.cameraDist))
}

func (g *Graphics3D) CalculateLuminance(normal Vec3) float64{
	light := g.light.Dot(normal)
	if light < 0 {
		light = 0
	}
	return light
}

func GetLuminanceChar(luminance float64) rune {
	chars := []rune(".:-=+*#%@")
	index := int(luminance * float64(len(chars)-1))
	if index >= len(chars) {
		index = len(chars) - 1
	}
	return chars[index]
}

func interpolateZ(x, y float64, projected [4]math2d.Vec2, depths [4]float64) float64 {
	// TODO: barycentric coordinates
	return (depths[0] + depths[1] + depths[2] + depths[3]) / 4
}

func (g *Graphics3D) DrawFace(cube *Cube, face Face, char rune) {
	var projected[4] math2d.Vec2 
	var depths[4] float64

	for i := 0; i < len(face.vertices); i++ {
		projected[i] = cube.vertices[face.vertices[i]].toVec2(g.cameraDist)
		depths[i] = cube.vertices[face.vertices[i]].Z
	}

	minX, maxX := projected[0].X, projected[0].X
	minY, maxY := projected[0].Y, projected[0].Y
	for _, p := range projected[1:] {
		minX = math.Min(minX, p.X)
		maxX = math.Max(maxX, p.X)
		minY = math.Min(minY, p.Y)
		maxY = math.Max(maxY, p.Y)
	}	

	style := tcell.StyleDefault.Foreground(tcell.ColorGreen)

	w, h := (*g.screen).Size()
	cx := float64(w) / 2
	cy := float64(h) / 2

	for y := minY; y <= maxY; y += 0.5 {
		for x := minX; x <= maxX; x++ {
			if insideRectangle(math2d.Vec2{X: x,Y: y}, projected) {
				sx := int(math.Round(x + cx))
				sy := int(math.Round(y/2 + cy))
				z := interpolateZ(x, y, projected, depths)
				if g.zbuffer.Test(sx, sy, z) {
					(*g.screen).SetContent(sx, sy, char, nil, style)
				}
			}
		}
	}
}

func insideRectangle(pt math2d.Vec2, rect[4] math2d.Vec2) bool {
	sign := func(pt, a, b math2d.Vec2) float64 {
		return b.Sub(a).Cross(pt.Sub(a))
	}

	b1 := sign(pt, rect[0], rect[1]) < 0
	b2 := sign(pt, rect[1], rect[2]) < 0
	b3 := sign(pt, rect[2], rect[3]) < 0
	b4 := sign(pt, rect[3], rect[0]) < 0

	return b1 == b2 && b2 == b3 && b3 == b4
}

func (g *Graphics3D) DrawCube(c *Cube) {
	for _, face := range c.faces {
		normal := c.GetFaceNormal(face)
		luminance := g.CalculateLuminance(normal)
		char := GetLuminanceChar(luminance)
		g.DrawFace(c, face, char)
	}
}
