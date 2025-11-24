package main

import (
	// "fmt"
	"math"
	"time"

	"github.com/gdamore/tcell/v2"
)

func rotate(x float64, y float64, t float64) (float64, float64) {
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

func draw_line(x1, y1, x2, y2 float64, screen tcell.Screen) {
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

func draw_rectangle(coords [4][2]float64, screen tcell.Screen, cx float64, cy float64) {
	for i := range len(coords) - 1 {
		// x := coords[i][0]
		// y := coords[i][1]
		// screen.SetContent(int(x)+cx, int(y*0.5)+cy, '@', nil, style)
		draw_line(coords[i][0]+cx, (coords[i][1]/2.0)+cy, coords[i+1][0]+cx, (coords[i+1][1]/2.0)+cy, screen)
	}
	draw_line(coords[3][0]+cx, (coords[3][1]/2.0)+cy, coords[0][0]+cx, (coords[0][1]/2.0)+cy, screen)
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()

	keys := make(chan *tcell.EventKey, 10)

	go func() {
		for {
			event := screen.PollEvent()
			switch ev := event.(type) {
			case *tcell.EventKey:
				keys <- ev
			}
		}
	}()

	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()

	angle := 0.05

	w, h := screen.Size()

	cx, cy := float64(w)/2.0, float64(h)/2.0

	rectangleWidth := 25
	rectangleHeight := 25

	offcenter := 15.0

	coords := [4][2]float64{
		{offcenter, 0},
		{float64(rectangleWidth) + offcenter, 0},
		{offcenter + float64(rectangleWidth), float64(rectangleHeight)},
		{offcenter, float64(rectangleHeight)},
	}

	for {
		select {
		case ev := <-keys:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			}
		case <-ticker.C:
			for i := range 4 {
				x := coords[i][0]
				y := coords[i][1]
				coords[i][0], coords[i][1] = rotate(x, y, angle)
			}

			screen.Clear()
			draw_rectangle(coords, screen, cx, cy)
			screen.Show()
		}
	}
}
