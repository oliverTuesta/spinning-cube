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

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()

	style := tcell.StyleDefault.Foreground(tcell.ColorGreen)

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

	angle := 0.0
	baseX, baseY := 30.0, 0.0
	w, h := screen.Size()
	cx, cy := w/2, h/2
	for {
		select {
		case ev := <-keys:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			}
		case <-ticker.C:
			x, y := rotate(baseX, baseY, angle)
			angle += 0.1

			screen.Clear()
			screen.SetContent(int(x)+cx, int(y*0.5)+cy, '@', nil, style)
			screen.Show()
		}
	}
}
