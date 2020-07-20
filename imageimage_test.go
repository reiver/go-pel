package pel_test

import (
	"github.com/reiver/go-pel"

	"image"

	"testing"
)

func TestImageImage_ColorLinkedIn(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var datum image.Image = pel.ColorLinked{}

	if nil == datum {
		t.Error("This should never happen")
	}
}

func TestImageImage_RGBA(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var datum image.Image = pel.RGBA{}

	if nil == datum {
		t.Error("This should never happen")
	}
}
