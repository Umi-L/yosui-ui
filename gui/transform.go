package gui

import "github.com/hajimehoshi/ebiten/v2"

type Transform struct {
	//absolute pixel values
	X float32
	Y float32
	W float32
	H float32

	//Relative % values
	XPercent float32
	YPercent float32
	WPercent float32
	HPercent float32

	Origin Origin
}

func MakeTransformWithImage(image *ebiten.Image, origin OriginIndex) Transform {
	return Transform{
		W: float32(image.Bounds().Max.X - image.Bounds().Min.X),
		H: float32(image.Bounds().Max.Y - image.Bounds().Min.Y),

		Origin: Origin{OriginIndex: origin},
	}
}

func (t Transform) SetHPercent(value float32) Transform {
	t.HPercent = value
	return t
}

func (t Transform) SetWPercent(value float32) Transform {
	t.WPercent = value
	return t
}

func (t Transform) SetXPercent(value float32) Transform {
	t.XPercent = value
	return t
}

func (t Transform) SetYPercent(value float32) Transform {
	t.YPercent = value
	return t
}

func (t Transform) SetH(value float32) Transform {
	t.H = value
	return t
}

func (t Transform) SetW(value float32) Transform {
	t.W = value
	return t
}

func (t Transform) SetX(value float32) Transform {
	t.X = value
	return t
}

func (t Transform) SetY(value float32) Transform {
	t.Y = value
	return t
}
