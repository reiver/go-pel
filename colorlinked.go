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

// At returns the color at an (x,y) coordinate.
//
// At helps ColorLinked  fit the Go built-in image.Image interface.
func (receiver ColorLinked) At(x, y int) color.Color {
	if receiver.X != x || receiver.Y != y {
		return RGBA{
			X: receiver.X,
			Y: receiver.Y,
		}
	}

	return receiver
}

// Bounds returns the (rectangular) area for which RGBA may have non-zero color.
//
// Bounds helps ColorLinked fit the Go built-in image.Image interface.
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

// ColorModel returns the color-model.
//
// ColorModel helps ColorLinked fit the Go built-in image.Image interface.
func (receiver ColorLinked) ColorModel() color.Model {
	return color.NRGBAModel
}

// RGBA returns the alpha-premultiplied red, green, blue and alpha values for the color.
//
// Bounds helps ColorLinked fit the Go built-in color.Color interface.
func (receiver ColorLinked) RGBA() (r, g, b, a uint32) {
	return receiver.Color.RGBA()
}
