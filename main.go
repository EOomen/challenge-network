package main

import (
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(runGame)
}
func runGame() {
	cfg := pixelgl.WindowConfig{
		Title:     ":)",
		Bounds:    pixel.R(0, 0, 1024, 750),
		VSync:     true,
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(false)

	load := asset.NewLoad(os.DirFS("./"))

	wolfSprite, err := load.Sprite("wolf_1.png")
	if err != nil {
		panic(err)
	}

	wolfPosition := win.Bounds().Center()

	for !win.JustPressed(pixelgl.KeyEscape) {
		win.Clear(pixel.RGB(255, 255, 255))

		//Take 'sprite name' and draw him to the window, scale/rotate him if needed and move him to the correct position
		wolfSprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.0).Moved(wolfPosition))

		if win.Pressed(pixelgl.KeyA) {
			wolfPosition.X -= 2.0
		}
		if win.Pressed(pixelgl.KeyD) {
			wolfPosition.X += 2.0
		}
		if win.Pressed(pixelgl.KeyS) {
			wolfPosition.Y -= 2.0
		}
		if win.Pressed(pixelgl.KeyW) {
			wolfPosition.Y += 2.0
		}

		win.Update()
	}
}
