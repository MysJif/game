package main;

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	conf := pixelgl.WindowConfig{
		Title: "Tetris",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync: true,
	}

	win, err := pixelgl.NewWindow(conf)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Skyblue)

	for !win.Closed() {
		win.Update()
	}
}
	
func main() {
	pixelgl.Run(run)
}