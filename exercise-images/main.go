package main

import(
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct{
	w, h int
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
