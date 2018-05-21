package main

import (
	"image"
	"os"

	_ "image/vector"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const IMAGE_DIR = "./assets/images"

func loadImage(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	conf := pixelgl.WindowConfig{
		Title:  "Tetris",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(conf)
	if err != nil {
		panic(err)
	}

	img, err := loadImage(IMAGE_DIR + "/pajaHandsUp.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(img, img.Bounds())

	win.Clear(colornames.Skyblue)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
