package main

import(
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct{
	w, h int
}

func (img Image) ColorModel() color.Mode {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
