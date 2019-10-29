package main

import (
	"fmt"
	"image"
	"image/color"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	r := color.RGBA{0x00, 0x77, 0x77, 0xff}
	m.SetRGBA(0, 0, r)
	fmt.Println(m.At(0, 0).RGBA())
}
