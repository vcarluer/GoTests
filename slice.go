package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    y := make([][]uint8, dy)
    for i:=0; i < dy; i++ {
        x := make([]uint8, dx)
        y[i] = x
        for j := 0; j < dx; j++ {
            x[j] = uint8((i+j) / 2)
        }
    }    
    
    return y
}

func main() {
    pic.Show(Pic)
}

