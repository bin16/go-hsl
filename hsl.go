package hsl

import (
	"fmt"
	"image/color"
	"math"
)

type HSL struct {
	H float64 // Hue        [0, 360)
	S float64 // Saturation [0, 1]
	L float64 // Lightness  [0, 1]
}

// to color.RGBA{}
func (hsl HSL) ToRGBA() color.RGBA {
	var (
		r0, g0, b0, a0 = hsl.RGBA()
	)

	return color.RGBA{
		uint8(r0 >> 8),
		uint8(g0 >> 8),
		uint8(b0 >> 8),
		uint8(a0 >> 8),
	}
}

func (clr HSL) RGBA() (r, g, b, a uint32) {
	var (
		h          = float64(clr.H)
		s          = float64(clr.S)
		l          = float64(clr.L)
		r0, g0, b0 = hslToRgb(h, s, l)
	)

	r = uint32(math.Round(r0 * 0xffff))
	g = uint32(math.Round(g0 * 0xffff))
	b = uint32(math.Round(b0 * 0xffff))
	a = 0xffff
	return
}

func (clr HSL) Equals(clr1 HSL) bool {
	return math.Round(clr.H) == math.Round(clr1.H) &&
		math.Round(clr.S*100) == math.Round(clr1.S*100) &&
		math.Round(clr.L*100) == math.Round(clr1.L*100)
}

func (clr HSL) String() string {
	return fmt.Sprintf(
		`HSL(%.1f, %.1f, %.1f)`,
		math.Round(clr.H),
		math.Round(clr.S*100),
		math.Round(clr.L*100),
	)
}
