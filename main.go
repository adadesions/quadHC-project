package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func cvtToGray(img image.Image) *image.Gray {
	canvas := image.NewGray(img.Bounds())
	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
			origin := img.At(i, j)
			canvas.Set(i, j, color.GrayModel.Convert(origin))
		}
	}

	return canvas
}

func saveToPNG(outputPath string, img *image.Gray) {
	outputFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Error at create output file: %v", err)
	}
	png.Encode(outputFile, img)
	outputFile.Close()
}

func info(img *image.Gray, isShowPix bool) {
	fmt.Println("============== Info ================")
	if isShowPix {
		fmt.Printf("Gray Data: %v", img)
	}
	fmt.Printf("Gray Bound: %v\n", img.Bounds())
	fmt.Printf("Gray Data: %v\n", img.Pix[0])
	fmt.Printf("Gray Len: %v\n", len(img.Pix))
	fmt.Println("====================================")
}

func ImRead(imgPath string) (*os.File, error) {
	img, err := os.Open(imgPath)

	if err != nil {
		return nil, err
	}

	return img, nil
}

func main() {
	img, err := ImRead("images/sample1.png")
	if err != nil {
		log.Fatalf("Error occured during reading an image file: %v", err)
	}
	defer img.Close()

	imgData, err := png.Decode(img)

	if err != nil {
		log.Fatalf("Error can't decode image: %v\n", err)
	}

	canvas := cvtToGray(imgData)
	saveToPNG("images/output.png", canvas)
	info(canvas, false)

}
