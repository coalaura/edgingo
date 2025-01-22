package edgingo

import "image"

type WalkerFunc func(*image.RGBA, CompareFunc, int, int, int) int

type CompareFunc func(int, int) bool

func WalkTopToBottom(rgba *image.RGBA, cmp CompareFunc, width, height, acceptable int) (edge int) {
	for y := 0; y < height; y++ {
		edge = y

		for x := 0; x <= width; x++ {
			if acceptable > 0 && (x >= acceptable && x <= width-acceptable) {
				continue
			}

			if !cmp(x, y) {
				return
			}
		}
	}

	return -1
}

func WalkBottomToTop(rgba *image.RGBA, cmp CompareFunc, width, height, acceptable int) (edge int) {
	for y := height; y > 0; y-- {
		edge = height - y

		for x := 0; x <= width; x++ {
			if acceptable > 0 && (x >= acceptable && x <= width-acceptable) {
				continue
			}

			if !cmp(x, y) {
				return
			}
		}
	}

	return -1
}

func WalkLeftToRight(rgba *image.RGBA, cmp CompareFunc, width, height, acceptable int) (edge int) {
	for x := 0; x < width; x++ {
		edge = x

		for y := 0; y <= height; y++ {
			if acceptable > 0 && (y >= acceptable && y <= height-acceptable) {
				continue
			}

			if !cmp(x, y) {
				return
			}
		}
	}

	return -1
}

func WalkRightToLeft(rgba *image.RGBA, cmp CompareFunc, width, height, acceptable int) (edge int) {
	for x := width; x > 0; x-- {
		edge = width - x

		for y := 0; y <= height; y++ {
			if acceptable > 0 && (y >= acceptable && y <= height-acceptable) {
				continue
			}

			if !cmp(x, y) {
				return
			}
		}
	}

	return -1
}
