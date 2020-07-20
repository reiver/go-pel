package pel_test

import (
	"github.com/reiver/go-pel"

	"image/color"

	"testing"
)

func TestColorColor_ColorLinked(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var datum color.Color = pel.ColorLinked{}

	if nil == datum {
		t.Error("This should never happen")
	}
}

func TestColorColor_RGBA(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var datum color.Color = pel.RGBA{}

	if nil == datum {
		t.Error("This should never happen")
	}
}
