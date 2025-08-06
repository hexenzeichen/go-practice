package main

import(
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
