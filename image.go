package main

import (
    "code.google.com/p/go-tour/pic"
    "image/color"
    "image"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
 	   return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    rect := image.Rect(0, 0, 512, 256)
    return rect
}

func (i Image) At(x, y int) color.Color {
    var v uint8 
    v = uint8((x + y) / 2)
    return color.RGBA{v, v, 255, 255}
}

func main() {
    m := Image{}
    pic.ShowImage(m)
}

