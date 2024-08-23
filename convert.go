package hsl

import (
	"image/color"
	"math"
)

// https://en.wikipedia.org/wiki/HSL_and_HSV#From_RGB
// r: [0, 1]
// g: [0, 1]
// b: [0, 1]
// h: [0, 360]
// s: [0, 1]
// l: [0, 1]
func rgbToHsl(r, g, b float64) (h, s, l float64) {
	var (
		max = math.Max(math.Max(r, g), b)
		min = math.Min(math.Min(r, g), b)
	)

	l = (min + max) / 2

	switch {
	case min == max:
		h = 0
	case max == r:
		h = 60 * math.Mod((g-b)/(max-min)+360, 6)
	case max == g:
		h = 60*(b-r)/(max-min) + 120
	case max == b:
		h = 60*(r-g)/(max-min) + 240
	}

	switch {
	case l == 0 || max == min:
		s = 0
	case l > 0 && l <= .5:
		s = (max - min) / (max + min)
	case l > .5:
		s = (max - min) / (2 - (max + min))
	}

	return
}

// https://en.wikipedia.org/wiki/HSL_and_HSV#HSL_to_RGB
// h: [0, 360]
// s: [0, 1]
// l: [0, 1]
// r: [0, 1]
// g: [0, 1]
// b: [0, 1]
func hslToRgb(h, s, l float64) (r, g, b float64) {
	var (
		c          = (1 - math.Abs(2*l-1)) * s
		h1         = h / 60
		x          = c * (1 - math.Abs(math.Mod(h1, 2)-1))
		r1, g1, b1 float64
		m          = l - c/2
	)

	switch {
	case h1 >= 0 && h1 < 1:
		r1, g1, b1 = c, x, 0
	case h1 >= 1 && h1 < 2:
		r1, g1, b1 = x, c, 0
	case h1 >= 2 && h1 < 3:
		r1, g1, b1 = 0, c, x
	case h1 >= 3 && h1 < 4:
		r1, g1, b1 = 0, x, c
	case h1 >= 4 && h1 < 5:
		r1, g1, b1 = x, 0, c
	case h1 >= 5 && h1 < 6:
		r1, g1, b1 = c, 0, x
	}

	r, g, b = r1+m, g1+m, b1+m
	return
}

func FromColor(clr color.Color) HSL {
	var (
		r0, g0, b0, _ = clr.RGBA()
		r             = float64(r0) / float64(0xffff)
		g             = float64(g0) / float64(0xffff)
		b             = float64(b0) / float64(0xffff)
		h, s, l       = rgbToHsl(r, g, b)
	)

	return HSL{
		h, s, l,
	}
}
