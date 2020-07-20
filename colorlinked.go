package pel

import (
	"image"
	"image/color"
)

type ColorLinked struct {
	X int
	Y int

	Color color.Color
}

func (receiver ColorLinked) At(x, y int) color.Color {
	if receiver.X != x || receiver.Y != y {
		return RGBA{
			X: receiver.X,
			Y: receiver.Y,
		}
	}

	return receiver
}

func (receiver ColorLinked) Bounds() image.Rectangle {
	const width  = 1
	const height = 1

	return image.Rectangle{
		Min: image.Point{
			X: receiver.X,
			Y: receiver.Y,
		},
		Max: image.Point{
			X: receiver.X+width,
			Y: receiver.Y+height,
		},
	}
}

func (receiver ColorLinked) ColorModel() color.Model {
	return color.NRGBAModel
}

func (receiver ColorLinked) RGBA() (r, g, b, a uint32) {
	return receiver.Color.RGBA()
}
