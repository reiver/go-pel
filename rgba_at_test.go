package pel_test

import (
	"github.com/reiver/go-pel"

	"image"
	"math"
	"math/rand"
	"time"

	"testing"
)

func TestRGBA_At_alpha255(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	for testNumber:=0; testNumber<10; testNumber++ {

		var x,y int
		{
			x = randomness.Int()
			if 0 == randomness.Int()%2 {
				x = -x
			}

			y = randomness.Int()
			if 0 == randomness.Int()%2 {
				y = -y
			}
		}

		var r,b,g,a uint8
		{
			r = uint8(randomness.Intn(256))
			g = uint8(randomness.Intn(256))
			b = uint8(randomness.Intn(256))
			a = 255
		}

		var pixel pel.RGBA = pel.RGBA{
			X:x,
			Y:y,

			R:r,
			G:g,
			B:b,
			A:a,
		}

		var img image.Image = pixel

		{
			eR := uint32(r) * (0xffff/0xff)
			eG := uint32(g) * (0xffff/0xff)
			eB := uint32(b) * (0xffff/0xff)
			eA := uint32(a) * (0xffff/0xff)

			aR, aG, aB, aA := img.At(x,y).RGBA()

			if eR != aR || eG != aG || eB != aB || eA != aA {
				t.Errorf("For test #%d, the actual color was not what was expected.", testNumber)
				t.Logf("(x,y)=(%d,%d)", x,y)
				t.Logf("rgba(%d,%d,%d,%d)", r,g,b,a)
				t.Logf("EXPECTED (r,g,b,a)=(%d,%d,%d,%d) ⤳ rgba(%d,%d,%d,%d)", eR, eG, eB, eA,  ((eR*0xff)/0xffff), ((eG*0xff)/0xffff), ((eB*0xff)/0xffff), ((eA*0xff)/0xffff))
				t.Logf("ACTUAL   (r,g,b,a)=(%d,%d,%d,%d) ⤳ rgba(%d,%d,%d,%d)", aR, aG, aB, aA,  ((aR*0xff)/0xffff), ((aG*0xff)/0xffff), ((aB*0xff)/0xffff), ((aA*0xff)/0xffff))
				continue
			}
		}

		for subTestNumber:=0; subTestNumber<30; subTestNumber++ {
			var xx,yy int
			{
				xx = randomness.Int()
				if 0 == randomness.Int()%2 {
					xx = -xx
				}

				yy = randomness.Int()
				if 0 == randomness.Int()%2 {
					yy = -yy
				}

				if x == xx && y == yy {
					xx--
					yy--
				}
			}

			var eR, eG, eB, eA uint32 = 0,0,0,0

			aR, aG, aB, aA := img.At(xx,yy).RGBA()

			if eR != aR || eG != aG || eB != aB || eA != aA {
				t.Errorf("For test #%d, the actual color was not what was expected.", testNumber)
				t.Logf("(x,y)=(%d,%d)", x,y)
				t.Logf("rgba(%d,%d,%d,%d)", r,g,b,a)
				t.Logf("EXPECTED (r,g,b,a)=(%d,%d,%d,%d) ⤳ rgba(%d,%d,%d,%d)", eR, eG, eB, eA,  ((eR*0xff)/0xffff), ((eG*0xff)/0xffff), ((eB*0xff)/0xffff), ((eA*0xff)/0xffff))
				t.Logf("ACTUAL   (r,g,b,a)=(%d,%d,%d,%d) ⤳ rgba(%d,%d,%d,%d)", aR, aG, aB, aA,  ((aR*0xff)/0xffff), ((aG*0xff)/0xffff), ((aB*0xff)/0xffff), ((aA*0xff)/0xffff))
				continue
			}
		}
	}
}

func TestRGBA_At_alpha127(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	for testNumber:=0; testNumber<10; testNumber++ {

		var x,y int
		{
			x = randomness.Int()
			if 0 == randomness.Int()%2 {
				x = -x
			}

			y = randomness.Int()
			if 0 == randomness.Int()%2 {
				y = -y
			}
		}

		var r,b,g,a uint8
		{
			r = uint8(randomness.Intn(256))
			g = uint8(randomness.Intn(256))
			b = uint8(randomness.Intn(256))
			a = 127
		}

		var pixel pel.RGBA = pel.RGBA{
			X:x,
			Y:y,

			R:r,
			G:g,
			B:b,
			A:a,
		}

		var img image.Image = pixel

		{
			eR := uint32(math.Round((float64(r) * 127.0) / 255.0))*(0xffff/0xff)
			eG := uint32(math.Round((float64(g) * 127.0) / 255.0))*(0xffff/0xff)
			eB := uint32(math.Round((float64(b) * 127.0) / 255.0))*(0xffff/0xff)
			eA := uint32(0xffff)

			aR, aG, aB, aA := img.At(x,y).RGBA()

			if eR != aR || eG != aG || eB != aB || eA != aA {
				t.Errorf("For test #%d, the actual color was not what was expected.", testNumber)
				t.Logf("(x,y)=(%d,%d)", x,y)
				t.Logf("rgba(%d,%d,%d,%d)", r,g,b,a)
				t.Logf("EXPECTED (r,g,b,a)=(%d,%d,%d,%d) ⤳ rgba(%d,%d,%d,%d)", eR, eG, eB, eA,  ((eR*0xff)/0xffff), ((eG*0xff)/0xffff), ((eB*0xff)/0xffff), ((eA*0xff)/0xffff))
				t.Logf("ACTUAL   (r,g,b,a)=(%d,%d,%d,%d) ⤳ rgba(%d,%d,%d,%d)", aR, aG, aB, aA,  ((aR*0xff)/0xffff), ((aG*0xff)/0xffff), ((aB*0xff)/0xffff), ((aA*0xff)/0xffff))
				continue
			}
		}

		for subTestNumber:=0; subTestNumber<30; subTestNumber++ {
			var xx,yy int
			{
				xx = randomness.Int()
				if 0 == randomness.Int()%2 {
					xx = -xx
				}

				yy = randomness.Int()
				if 0 == randomness.Int()%2 {
					yy = -yy
				}

				if x == xx && y == yy {
					xx--
					yy--
				}
			}

			var eR, eG, eB, eA uint32 = 0,0,0,0

			aR, aG, aB, aA := img.At(xx,yy).RGBA()

			if eR != aR || eG != aG || eB != aB || eA != aA {
				t.Errorf("For test #%d, the actual color was not what was expected.", testNumber)
				t.Logf("(x,y)=(%d,%d)", x,y)
				t.Logf("rgba(%d,%d,%d,%d)", r,g,b,a)
				t.Logf("EXPECTED (r,g,b,a)=(%d,%d,%d,%d) ⤳ rgba(%d,%d,%d,%d)", eR, eG, eB, eA,  ((eR*0xff)/0xffff), ((eG*0xff)/0xffff), ((eB*0xff)/0xffff), ((eA*0xff)/0xffff))
				t.Logf("ACTUAL   (r,g,b,a)=(%d,%d,%d,%d) ⤳ rgba(%d,%d,%d,%d)", aR, aG, aB, aA,  ((aR*0xff)/0xffff), ((aG*0xff)/0xffff), ((aB*0xff)/0xffff), ((aA*0xff)/0xffff))
				continue
			}
		}
	}
}
