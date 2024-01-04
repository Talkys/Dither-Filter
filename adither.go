package main

import (
	"math"
)

func dither(inputMatrix float64, x, y, c, pattern, levels int) float64 {
	var mask float64

	if pattern == 1 {
		mask = float64((x^y*149)*1234&511) / 511.0
	} else if pattern == 2 {
		mask = float64(((x+c*17)^y*149)*1234&511) / 511.0
	} else if pattern == 3 {
		mask = float64((x+y*237)*119&255) / 255.0
	} else if pattern == 4 {
		mask = float64(((x+c*67)+y*236)*119&255) / 255.0
	} else if pattern == 5 {
		mask = 0.5
	} else {
		return inputMatrix
	}

	return math.Floor(float64(levels)*inputMatrix+mask) / float64(levels)
}

func adither(imageMatrix [][][3]uint8, pattern, levelR, levelG, levelB int) [][][3]uint8 {
	for i := range imageMatrix {
		for j := range imageMatrix[i] {
			c := 0
			imageMatrix[i][j][c] = uint8(dither(float64(imageMatrix[i][j][c])/255.0, j, i, c, pattern, levelR) * 255.0)
			c = 1
			imageMatrix[i][j][c] = uint8(dither(float64(imageMatrix[i][j][c])/255.0, j, i, c, pattern, levelG) * 255.0)
			c = 2
			imageMatrix[i][j][c] = uint8(dither(float64(imageMatrix[i][j][c])/255.0, j, i, c, pattern, levelB) * 255.0)
		}
	}
	return imageMatrix
}
