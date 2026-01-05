package math3d

import "math"

type ZBuffer struct {
	width, height int
	depths [][]float64
}

func NewZBuffer(w, h int) *ZBuffer {
	depths := make([][]float64, h)
	for i := range depths {
		depths[i] = make([]float64, w)
		for j := range depths[i] {
			depths[i][j] = math.Inf(1)
		}
	}
	return &ZBuffer{depths: depths, width: w, height: h}
}

func (zb *ZBuffer) Clear() {
	for i := range zb.depths {
		for j := range zb.depths[i] {
			zb.depths[i][j] = math.Inf(1)
		}
	}
}

func (zb *ZBuffer) Test(x, y int, z float64) bool {
	if x < 0 || x >= zb.width || y < 0 || y >= zb.height {
		return false
	}
	if z < zb.depths[y][x] {
		zb.depths[y][x] = z
		return true
	}
	return false
}
