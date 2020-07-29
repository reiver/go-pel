/*
Package pel provides a type that represents a single pel — that represents a single pixel — at a specific (x,y) location,
that seamlessly works with Go's built-in "image", "image/color", and "image/draw" packages.

A ‘pel’ is also often called a ‘pixel’.

Example

Package pel provides 2 different way of creating pels — of creating pixels:

Here is the 1st way:

	import "github.com/reiver/go-pel"
	
	// ...
	
	pixel := pel.RGBA{
		
		// (x,y) = (95,73)
		X:95,
		Y:73,
		
		// rgba(255,199,6,255)
		R:255,
		G:199,
		B:6,
		A:255,
	}

And here is the 2nd way:

	import "github.com/reiver/go-pel"
	
	// ...
	
	var c color.Color
	
	// ...
	
	pixel := pel.ColorLinked{
		
		// (x,y) = (95,73)
		X:95,
		Y:73,
		
		Color: c,
	}

In both cases `pixel` fits he image.Image interface, and you can use it to draw. For example:

	draw.Draw(dst, pixel.Bounds(), pixel. pixel.Bound().Min, draw.Over)
*/
package pel
