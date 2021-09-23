package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func cvtToGray(img image.Image) *image.RGBA {
	canvas := image.NewRGBA(img.Bounds())
	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
			origin := img.At(i, j)
			canvas.Set(i, j, color.GrayModel.Convert(origin))
		}
	}

	return canvas
}

func saveToPNG(outputPath string, img *image.RGBA) {
	outputFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Error at create output file: %v", err)
	}
	png.Encode(outputFile, img)
	outputFile.Close()
}

func main() {
	img, err := os.Open("images/sample1.png")
	if err != nil {
		log.Fatalf("Error can't open image: %v\n", err)
	}

	defer img.Close()

	imgData, err := png.Decode(img)

	if err != nil {
		log.Fatalf("Error can't decode image: %v\n", err)
	}

	fmt.Printf("Bounds: %v\n", imgData.Bounds())
	fmt.Printf("ColorModel: %v\n", imgData.ColorModel())
	fmt.Printf("At(0,0): %v\n", imgData.At(0, 0))

	canvas := cvtToGray(imgData)
	saveToPNG("images/output.png", canvas)
}
