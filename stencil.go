package edgingo

import (
	"image"
	"image/draw"
)

type Stencil struct {
	Top    int
	Bottom int
	Left   int
	Right  int
}

func NewStencil() *Stencil {
	return &Stencil{
		Top:    0,
		Bottom: 0,
		Left:   0,
		Right:  0,
	}
}

func (s *Stencil) Set(side ImageSide, value int) {
	switch side {
	case SideTop:
		s.Top = value
	case SideBottom:
		s.Bottom = value
	case SideLeft:
		s.Left = value
	case SideRight:
		s.Right = value
	}
}

func (s *Stencil) Cut(rgba *image.RGBA, width, height int) *image.RGBA {
	if s.Top == 0 && s.Bottom == 0 && s.Left == 0 && s.Right == 0 {
		return rgba
	}

	cropped := image.NewRGBA(image.Rect(0, 0, width-(s.Left+s.Right), height-(s.Top+s.Bottom)))
	offset := image.Pt(s.Left, s.Top)

	draw.Draw(cropped, cropped.Bounds(), rgba, offset, draw.Src)

	return cropped
}
