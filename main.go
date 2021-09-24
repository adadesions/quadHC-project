package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/webview/webview"
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

func ImReadPNG(imgPath string) (image.Image, error) {
	imgFile, err := os.Open(imgPath)

	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	img, err := png.Decode(imgFile)

	return img, nil
}

func display(imgPath string) {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	baseURL := "http://localhost"
	port := "8100"

	w.SetTitle("QuadHC Experimental")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(fmt.Sprintf("%s:%s/images/%s", baseURL, port, imgPath))
	w.Run()
}

func startServer() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func main() {
	go startServer()
	// img, err := ImReadPNG("images/sample1.png")

	// imgFile, err := os.Open("data/kaggle_3m/TCGA_CS_4941_19960909/TCGA_CS_4941_19960909_1.tif")
	imgFile, err := os.Open("images/sample1.png")
	if err != nil {
		log.Fatalf("Error occured during reading an image file: %v", err)
	}
	defer imgFile.Close()
	imgFile.Seek(0, 0)

	imgTiff, extension, err := image.Decode(imgFile)

	if err != nil {
		log.Fatalf("Error can't decode image: %v\n", err)
	}

	fmt.Printf("Img extension: %v\n", extension)

	canvas := cvtToGray(imgTiff)
	saveToPNG("images/output.png", canvas)
	info(canvas, false)
	display("output.png")
}
