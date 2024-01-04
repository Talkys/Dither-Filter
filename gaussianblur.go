package main

import (
	"math"
	"sync"
)

func gaussianblur(imageMatrix [][][3]uint8, sigma float64) [][][3]uint8 {
	height := len(imageMatrix)
	width := len(imageMatrix[0])

	// Calcula o tamanho do kernel
	size := int(6*sigma + 1)
	if size%2 == 0 {
		size++
	}

	// Cria o kernel gaussiano
	kernel := make([][]float64, size)
	for i := range kernel {
		kernel[i] = make([]float64, size)
		for j := range kernel[i] {
			x := float64(i-size/2) / sigma
			y := float64(j-size/2) / sigma
			kernel[i][j] = math.Exp(-(x*x + y*y) / 2)
		}
	}

	// Normaliza o kernel
	sum := 0.0
	for i := range kernel {
		for j := range kernel[i] {
			sum += kernel[i][j]
		}
	}
	for i := range kernel {
		for j := range kernel[i] {
			kernel[i][j] /= sum
		}
	}

	var wg sync.WaitGroup

	apply_blur := func(c int) {
		defer wg.Done()
		tempMatrix := make([][]float64, height)
		for i := range tempMatrix {
			tempMatrix[i] = make([]float64, width)
		}

		//Magia negra
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				pixelValue := 0.0
				for i := 0; i < size; i++ {
					for j := 0; j < size; j++ {
						xx := x - size/2 + i
						yy := y - size/2 + j
						if xx >= 0 && xx < width && yy >= 0 && yy < height {
							pixelValue += float64(imageMatrix[yy][xx][c]) * kernel[i][j]
						}
					}
				}
				tempMatrix[y][x] = pixelValue
			}
		}

		// Atualiza a matriz original com os valores borrados
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				imageMatrix[y][x][c] = uint8(tempMatrix[y][x])
			}
		}
	}

	// Aplica o blur gaussiano à matriz
	for c := 0; c < 3; c++ {
		wg.Add(1)
		go apply_blur(c)
	}

	wg.Wait()

	return imageMatrix
}
