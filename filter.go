package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"os"
)

func loadImage(filePath string) ([][][3]uint8, error) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	imgArray := make([][][3]uint8, height)
	for i := range imgArray {
		imgArray[i] = make([][3]uint8, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			r, g, b, _ := img.At(j, i).RGBA()
			imgArray[i][j] = [3]uint8{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
		}
	}

	return imgArray, nil
}

func saveImage(imgArray [][][3]uint8, outputPath string) error {
	height, width := len(imgArray), len(imgArray[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			img.Set(j, i, color.RGBA{imgArray[i][j][0], imgArray[i][j][1], imgArray[i][j][2], 255})
		}
	}

	imgFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	err = png.Encode(imgFile, img)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var inputImagePath, outputImagePath, paletteFilePath string

	// Definindo flags
	flag.StringVar(&inputImagePath, "i", "", "Input image path")
	flag.StringVar(&outputImagePath, "o", "", "Output image path")
	flag.StringVar(&paletteFilePath, "p", "", "Palette file path")

	// Parse os argumentos da linha de comando
	flag.Parse()

	// Verifica se os argumentos obrigatórios foram fornecidos
	if inputImagePath == "" || outputImagePath == "" || paletteFilePath == "" {
		fmt.Println("Usage: -i <inputImagePath> -o <outputImagePath> -p <paletteFilePath>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("Loading image")

	imageMatrix, err := loadImage(inputImagePath)
	if err != nil {
		fmt.Println("Error loading image:", err)
		return
	}

	fmt.Println("Image loaded")
	fmt.Println("Applying dithering")

	imageMatrix = adither(imageMatrix, 4, 5, 4, 5)

	fmt.Println("Dithering done")
	fmt.Println("Mapping palette")

	imageMatrix, err = changePalette(imageMatrix, paletteFilePath)
	if err != nil {
		fmt.Println("Error changing palette:", err)
		return
	}

	fmt.Println("Colors changed")
	fmt.Println("Applying gaussian blur")

	imageMatrix = gaussianblur(imageMatrix, 0.60)

	fmt.Println("Blur done")
	fmt.Println("Saving image")

	err = saveImage(imageMatrix, outputImagePath)
	if err != nil {
		fmt.Println("Error saving image:", err)
	}

	fmt.Println("Image saved successfully")
}
