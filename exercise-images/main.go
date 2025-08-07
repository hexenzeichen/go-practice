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

func main() {
	m := Image{}
	pic.ShowImage(m)
}
