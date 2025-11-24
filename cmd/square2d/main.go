package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/olivertuesta/spinning-donut/internal/math2d"
)

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
				coords[i][0], coords[i][1] = math2d.Rotate(x, y, angle)
			}

			screen.Clear()
			math2d.DrawRectangle(coords, screen, cx, cy)
			screen.Show()
		}
	}
}
