package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (p Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (p Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (p Image) At(x, y int) color.Color {
	v := uint8(x ^ y - y ^ x)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
