package edgingo

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestEdge(t *testing.T) {
	file, err := os.OpenFile("test.png", os.O_RDONLY, 0)
	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	rgba := AsRGBA(img)

	assertWidth(t, RemoveHorizontalEdges(rgba, false), 488)
	assertWidth(t, RemoveHorizontalEdges(rgba, true), 446)
	assertHeight(t, RemoveVerticalEdges(rgba, false), 484)
	assertHeight(t, RemoveVerticalEdges(rgba, true), 446)

	rgba = RemoveAllEdges(rgba, true)

	green := color.RGBA{0, 255, 0, 255}

	assertColor(t, rgba, 0, 0, green)
	assertColor(t, rgba, 445, 445, green)

	output, err := os.OpenFile("output.png", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatal(err)
	}

	defer output.Close()

	err = png.Encode(output, rgba)
	if err != nil {
		t.Fatal(err)
	}
}

func assertColor(t *testing.T, rgba *image.RGBA, x, y int, expected color.Color) {
	actual := rgba.At(x, y)

	r1, g1, b1, _ := actual.RGBA()
	r2, g2, b2, _ := expected.RGBA()

	if r1 != r2 || g1 != g2 || b1 != b2 {
		t.Errorf("expected color %v, got %v", expected, actual)
	}
}

func assertWidth(t *testing.T, img *image.RGBA, width int) {
	actual := img.Bounds().Dx()

	if actual != width {
		t.Errorf("expected width %d, got %d", width, actual)
	}
}

func assertHeight(t *testing.T, img *image.RGBA, height int) {
	actual := img.Bounds().Dy()

	if actual != height {
		t.Errorf("expected height %d, got %d", height, actual)
	}
}
