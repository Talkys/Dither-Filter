package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func changePalette(imageMatrix [][][3]uint8, paletteFilePath string) ([][][3]uint8, error) {
	// Carrega a paleta do arquivo
	palette, err := loadPalette(paletteFilePath)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	mapx := func(y int) {
		defer wg.Done()
		for x := 0; x < len(imageMatrix[y]); x++ {
			imageMatrix[y][x] = findClosestColorInPalette(imageMatrix[y][x], palette)
		}
	}

	// Mapeia cada cor na imagem para a cor mais próxima na paleta
	for y := 0; y < len(imageMatrix); y++ {
		wg.Add(1)
		go mapx(y)
	}

	wg.Wait()

	return imageMatrix, nil
}

func loadPalette(filePath string) ([][3]uint8, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var palette [][3]uint8
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hexCode := scanner.Text()
		colorRGB, err := hexToRGB(hexCode)
		if err != nil {
			return nil, err
		}
		palette = append(palette, colorRGB)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return palette, nil
}

func hexToRGB(hexCode string) ([3]uint8, error) {
	var rgb [3]uint8
	_, err := fmt.Sscanf(hexCode, "%02x%02x%02x", &rgb[0], &rgb[1], &rgb[2])
	if err != nil {
		return rgb, err
	}
	return rgb, nil
}

func findClosestColorInPalette(pixel [3]uint8, palette [][3]uint8) [3]uint8 {
	var closestColor [3]uint8
	minDistance := uint32(^uint32(0)) // Valor máximo possível de uint32

	for _, colorRGB := range palette {
		distance := colorDistance(pixel, colorRGB)
		if distance < minDistance {
			minDistance = distance
			closestColor = colorRGB
		}
	}

	return closestColor
}

func colorDistance(c1, c2 [3]uint8) uint32 {
	var distance uint32
	for i := 0; i < 3; i++ {
		d := uint32(c1[i]) - uint32(c2[i])
		distance += d * d
	}
	return distance
}
