package edgingo

import (
	"image"
	"image/color"
)

type ImageSide int

const (
	SideTop ImageSide = iota
	SideBottom
	SideLeft
	SideRight
)

func RemoveAllEdges(rgba *image.RGBA, aggressive bool) *image.RGBA {
	return RemoveEdges(rgba, []ImageSide{
		SideTop,
		SideBottom,
		SideLeft,
		SideRight,
	}, aggressive)
}

func RemoveVerticalEdges(rgba *image.RGBA, aggressive bool) *image.RGBA {
	return RemoveEdges(rgba, []ImageSide{
		SideTop,
		SideBottom,
	}, aggressive)
}

func RemoveHorizontalEdges(rgba *image.RGBA, aggressive bool) *image.RGBA {
	return RemoveEdges(rgba, []ImageSide{
		SideLeft,
		SideRight,
	}, aggressive)
}

func RemoveEdge(rgba *image.RGBA, side ImageSide, aggressive bool) *image.RGBA {
	return RemoveEdges(rgba, []ImageSide{side}, aggressive)
}

func RemoveEdges(rgba *image.RGBA, sides []ImageSide, aggressive bool) *image.RGBA {
	var (
		bounds = rgba.Bounds()
		width  = bounds.Dx()
		height = bounds.Dy()

		stencil = NewStencil()

		acceptable int
		amount     int
	)

	for _, side := range sides {
		if aggressive {
			if side == SideTop || side == SideBottom {
				acceptable = height / 10
			} else {
				acceptable = width / 10
			}
		}

		amount = edge(rgba, width, height, side, acceptable)

		stencil.Set(side, amount)
	}

	return stencil.Cut(rgba, width, height)
}

func edge(rgba *image.RGBA, width, height int, side ImageSide, acceptable int) int {
	var (
		test    color.Color
		against color.Color
		walker  WalkerFunc
	)

	switch side {
	case SideTop:
		walker = WalkTopToBottom
	case SideBottom:
		walker = WalkBottomToTop
	case SideLeft:
		walker = WalkLeftToRight
	case SideRight:
		walker = WalkRightToLeft
	}

	return walker(rgba, func(x, y int) bool {
		against = rgba.At(x, y)

		if test == nil {
			test = against

			return true
		}

		return equals(test, against)
	}, width-1, height-1, acceptable)
}

func equals(color1, color2 color.Color) bool {
	r1, g1, b1, _ := color1.RGBA()
	r2, g2, b2, _ := color2.RGBA()

	return int(r1)-int(r2) < 4 && int(g1)-int(g2) < 4 && int(b1)-int(b2) < 4
}
