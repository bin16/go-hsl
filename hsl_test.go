package hsl

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.rapidtables.com/convert/color/rgb-to-hsl.html
var colors = map[color.RGBA]HSL{
	{0x00, 0x00, 0x00, 0xff}: {0, 0, 0},        // Black
	{0xff, 0xff, 0xff, 0xff}: {0, 0, 1},        // White
	{0xff, 0x00, 0x00, 0xff}: {0, 1.0, 0.5},    // Red
	{0x00, 0xff, 0x00, 0xff}: {120, 1, 0.5},    // Lime
	{0x00, 0x00, 0xff, 0xff}: {240, 1.0, 0.5},  // Blue
	{0xff, 0xff, 0x00, 0xff}: {60, 1.0, 0.5},   // Yellow
	{0x00, 0xff, 0xff, 0xff}: {180, 1.0, 0.5},  // Cyan
	{0xff, 0x00, 0xff, 0xff}: {300, 1.0, 0.5},  // Magenta
	{0xbf, 0xbf, 0xbf, 0xff}: {0, 0, 0.75},     // Silver
	{0x80, 0x80, 0x80, 0xff}: {0, 0, 0.5},      // Gray
	{0x80, 0x00, 0x00, 0xff}: {0, 1.0, 0.25},   // Maroon
	{0x80, 0x80, 0x00, 0xff}: {60, 1.0, 0.25},  // Olive
	{0x00, 0x80, 0x00, 0xff}: {120, 1.0, 0.25}, // Green
	{0x80, 0x00, 0x80, 0xff}: {300, 1.0, 0.25}, // Purple
	{0x00, 0x80, 0x80, 0xff}: {180, 1.0, 0.25}, // Teal
	{0x00, 0x00, 0x80, 0xff}: {240, 1.0, 0.25}, // Navy

	{200, 200, 20, 0xFF}:  {60.00, .8182, .4313},
	{201, 200, 20, 0xFF}:  {59.67, .8190, .4333},
	{155, 245, 135, 0xFF}: {109.09, .8462, .7451},
}

func TestHsl(t *testing.T) {
	for c0 := range colors {
		var (
			hsl = FromColor(c0)
			c1  = hsl.ToRGBA()
		)

		assert.Equal(t, c0, c1)
	}
}

func TestToHSL(t *testing.T) {
	for d, c0 := range colors {
		c1 := FromColor(d)
		assert.True(t, c1.Equals(c0), "%v should be %v", c1, c0)
	}
}

func TestToRGB(t *testing.T) {
	for c0, d := range colors {
		c1 := d.ToRGBA()
		assert.Equal(t, c0, c1)
	}
}
