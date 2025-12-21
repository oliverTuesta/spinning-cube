package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/olivertuesta/spinning-cube/internal/math3d"
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

	g3 := math3d.NewGraphics3D(&screen, 100)
	cube := math3d.NewCube(15)
	cube.Move(10, 5, 3)

	for {
		select {
		case ev := <-keys:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			}
		case <-ticker.C:

			cube.RotateY(0.04)
			cube.RotateX(0.005)

			screen.Clear()
			g3.DrawCube(&cube)
			screen.Show()
		}
	}
}

