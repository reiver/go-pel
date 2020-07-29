package pel

import (
	"image"
	"image/color"
)

// RGBA represents a pel — a pixel — at a specific (x,y) location, and with a specific RGBA color.
//
// Here is an example of how to use it:
//
//	import "github.com/reiver/go-pel"
//	
//	// ...
//	
//	pixel := pel.RGBA{
//		
//		// (x,y) = (95,73)
//		X:95,
//		Y:73,
//		
//		// rgba(255,199,6,255)
//		R:255,
//		G:199,
//		B:6,
//		A:255,
//	}
//
// RGBA fits both the Go built-in image.Image & color.Color interfaces.
type RGBA struct {
	X int
	Y int

	R uint8
	G uint8
	B uint8
	A uint8
}

// At returns the color at an (x,y) coordinate.
//
// At helps RGBA fit the Go built-in image.Image interface.
func (receiver RGBA) At(x, y int) color.Color {
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
// Bounds helps RGBA fit the Go built-in image.Image interface.
func (receiver RGBA) Bounds() image.Rectangle {
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
// ColorModel helps RGBA fit the Go built-in image.Image interface.
func (receiver RGBA) ColorModel() color.Model {
	return color.NRGBAModel
}

// RGBA returns the alpha-premultiplied red, green, blue and alpha values for the color.
//
// Bounds helps RGBA fit the Go built-in color.Color interface.
func (receiver RGBA) RGBA() (r, g, b, a uint32) {
	r = uint32(receiver.R)
	g = uint32(receiver.G)
	b = uint32(receiver.B)
	a = uint32(receiver.A)

	if 0 == a {
		return 0,0,0,0
	}
	if 255 == a {
		r *= (0xffff/0xff)
		g *= (0xffff/0xff)
		b *= (0xffff/0xff)
		a *= (0xffff/0xff)

		return r,g,b,a
	}

	ra := r*a
	ga := g*a
	ba := b*a

	ramod := ra % 0xff
	gamod := ga % 0xff
	bamod := ba % 0xff

	r = ra / 0xff
	g = ga / 0xff
	b = ba / 0xff
	a = 0xff

	if 127 < ramod {
		r++
	}
	if 127 < gamod {
		g++
	}
	if 127 < bamod {
		b++
	}

	r *= (0xffff/0xff)
	g *= (0xffff/0xff)
	b *= (0xffff/0xff)
	a *= (0xffff/0xff)


	return r,g,b,a
}
