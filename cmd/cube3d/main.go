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

	g3 := math3d.NewGraphics3D(&screen, 1000)
	cubeWidth := 30.0
	cube := math3d.NewCube(cubeWidth)
	cube.Move(-cubeWidth/2,-cubeWidth/2,500)

	for {
		select {
		case ev := <-keys:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			}
		case <-ticker.C:

			cube.RotateAroundCenter(0.05,0.02,0)

			screen.Clear()
			g3.DrawCube(&cube)
			screen.Show()
		}
	}
}

