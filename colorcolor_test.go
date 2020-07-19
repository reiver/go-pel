package pel_test

import (
	"github.com/reiver/go-pel"

	"image/color"

	"testing"
)

func TestColorColor(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var datum color.Color = pel.RGBA{}

	if nil == datum {
		t.Error("This should never happen")
	}
}
