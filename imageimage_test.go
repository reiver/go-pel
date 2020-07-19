package pel_test

import (
	"github.com/reiver/go-pel"

	"image"

	"testing"
)

func TestImageImage(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var datum image.Image = pel.RGBA{}

	if nil == datum {
		t.Error("This should never happen")
	}
}
