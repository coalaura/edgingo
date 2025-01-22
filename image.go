package edgingo

import (
	"image"
	"image/draw"
)

// AsRGBA converts an image to an RGBA image
func AsRGBA(img image.Image) *image.RGBA {
	rgba, ok := img.(*image.RGBA)
	if ok {
		return rgba
	}

	rgba = image.NewRGBA(img.Bounds())

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)

	return rgba
}
