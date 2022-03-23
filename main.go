package main

import (
	"image"
	"image/draw"
	"image/png"
	_ "image/png"
	"os"
)

func getImage(filename string) image.Image {
	f, _ := os.Open(filename + ".png")
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return img
}

func main() {
	sampleImg := getImage("images/screenshot/sample")
	sampleMask := getImage("images/mask/sample")

	out := image.NewRGBA(sampleImg.Bounds())
	draw.DrawMask(out, out.Bounds(), sampleImg, image.Point{0, 0}, sampleMask, image.Point{0, 0}, draw.Over)

	file_out, _ := os.Create("sample.png")
	defer file_out.Close()
	png.Encode(file_out, out)
}
