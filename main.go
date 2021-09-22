package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

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

	// Modify Image
	canvas := image.NewRGBA(imgData.Bounds())
	fmt.Printf("Canvas Stride: %v\n", canvas.Stride)
	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
		// 	origin := imgData.At(i, j)
		// 	r, g, b, a := grayImg.RGBA()

			canvas.Set(i, j, color.GrayModel.Convert(imgData.At(i, j)))

		}
	}

	outputFile, err := os.Create("images/output.png")
	if err != nil {
		log.Fatalf("Error at create output file: %v", err)
	}

	png.Encode(outputFile, canvas)
	outputFile.Close()

}
