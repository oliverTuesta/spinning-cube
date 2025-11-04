package main

import (
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

func draw_rectangle(coords [4][2]float64, screen tcell.Screen, cx int, cy int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorGreen)
	for i := 0; i < 4; i++ {
		x := coords[i][0]
		y := coords[i][1]
		screen.SetContent(int(x)+cx, int(y*0.5)+cy, '@', nil, style)
	}
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

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	angle := 0.1

	w, h := screen.Size()

	cx, cy := w/2, h/2

	rectangleWidth := 10
	rectangleHeight := 15

	coords := [4][2]float64{
		{0, 0},
		{float64(rectangleWidth), 0},
		{float64(rectangleWidth), float64(rectangleHeight)},
		{0, float64(rectangleHeight)},
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
