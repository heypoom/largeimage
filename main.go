package main

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/dustin/go-humanize"
	"golang.org/x/image/colornames"
	"gopkg.in/cheggaaa/pb.v1"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
	"strconv"
)

const (
	JPEG = iota
	PNG
)

func generateFancy(width, height, format int) {
	fmt.Println("Initializing...")

	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	fmt.Println("Setting image buffer...")

	bar := pb.StartNew(width)

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			randU := func() uint8 {
				return uint8(rand.Intn(255))
			}

			c := color.RGBA{R: randU(), G: randU(), B: randU(), A: 255}
			img.Set(x, y, c)

		}

		bar.Increment()
	}

	bar.FinishPrint("Buffer Created!")

	fmt.Println("Creating Image...")

	var ext string

	switch format {
	case PNG: ext = "png"
	case JPEG: ext = "jpg"
	}

	fileName := "image." + ext

	// Encode as PNG.
	f, _ := os.Create(fileName)

	var err error

	switch format {
	case PNG:
		err = png.Encode(f, img)
	case JPEG:
		err = jpeg.Encode(f, img, nil)
	}

	if err != nil {
		fmt.Println("Unable to create the image...")
	}

	size := getFileSize(f)

	fmt.Println("Done! Image saved to:", fileName, "| size:", size)
}

func getFileSize(file *os.File) string {
	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Println("Cannot stat file...")
		return ""
	}

	return humanize.Bytes(uint64(fileInfo.Size()))
}

func _generateSimple(width, height int) {
	c := colornames.Salmon

	img := imaging.New(width, height, c)

	fmt.Println("Saving Buffers..")

	fileName := "image.jpg"

	err := imaging.Save(img, fileName)
	if err != nil {
		fmt.Println("Unable to save the file.")
	}

	fmt.Println("Done! Image saved to:", fileName)
}

func main() {
	size := 1000

	format := JPEG

	argc := len(os.Args)

	if argc > 1 {
		s, err := strconv.Atoi(os.Args[1])
		size = s

		if err != nil {
			fmt.Println("The size is not an integer!")
			return
		}
	}

	if argc > 2 {
		formatFlag := os.Args[2]
		fmt.Println("File format:", formatFlag)

		switch formatFlag {
		case "png":
			format = PNG
		case "jpeg":
			format = JPEG
		}
	}

	if len(os.Args) > 3 {
		fmt.Println("Usage: largeimage <size>")
		return
	}

	generateFancy(size, size, format)
}