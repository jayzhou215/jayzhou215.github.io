package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var slice [][]uint8
	for y := 0; y < dy; y++ {
		var yslice []uint8
		for x := 0; x < dx; x++ {
			yslice = append(yslice, uint8((x+y)/2))
		}
		slice = append(slice, yslice)
	}
	return slice
}

func main() {
	pic.Show(Pic)
}
