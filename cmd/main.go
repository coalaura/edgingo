package main

import (
	"image"
	"os"
	"strings"

	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"

	"github.com/coalaura/edgingo"
	"github.com/gen2brain/webp"
)

func main() {
	if len(os.Args) < 3 {
		println("Usage: edgingo <image> <output>")

		return
	}

	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0)
	if err != nil {
		println("Error opening file:", err)

		return
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		println("Error decoding image:", err)

		return
	}

	path := os.Args[2]

	output, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		println("Error opening output file:", err)

		return
	}

	defer output.Close()

	rgba := edgingo.RemoveAllEdges(edgingo.AsRGBA(img), true)

	if strings.HasSuffix(path, ".png") {
		err = png.Encode(output, rgba)
	} else if strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") {
		err = jpeg.Encode(output, rgba, &jpeg.Options{
			Quality: 90,
		})
	} else {
		err = webp.Encode(output, rgba, webp.Options{
			Method:   6,
			Lossless: true,
		})
	}

	if err != nil {
		println("Error encoding image:", err)

		return
	}

	println("Removed edges successfully!")
}
