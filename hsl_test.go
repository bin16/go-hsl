package hsl

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHsl(t *testing.T) {

	var colors = []color.RGBA{
		{39, 45, 45, 255},
		{0, 0, 160, 255},
		{0, 160, 0, 255},
		{160, 0, 0, 255},
		{0, 200, 160, 255},
		{200, 160, 0, 255},
		{160, 0, 200, 255},
		{50, 200, 160, 255},
		{200, 160, 30, 255},
		{160, 78, 200, 255},
	}

	for _, clr := range colors {
		var (
			hsl = FromColor(clr)
			c1  = hsl.ToRGBA()
		)

		assert.Equal(t, clr, c1)
	}
}

func TestToHSL(t *testing.T) {
	var colors = map[color.RGBA]HSL{
		{35, 206, 107, 255}:  {145, .71, .47},
		{39, 45, 45, 255}:    {180, .07, .16},
		{246, 248, 255, 255}: {227, 1, .98},
		{168, 70, 160, 255}:  {305, .41, .47},
		{80, 81, 79, 255}:    {90, .01, .31},
	}

	for d, c0 := range colors {
		c1 := FromColor(d)
		assert.True(t, c1.Equals(c0), "%v should be %v", c1, c0)
	}
}

func TestToRGB(t *testing.T) {
	var colors = map[color.RGBA]HSL{
		{35, 206, 107, 255}:  {145, .71, .47},
		{39, 45, 45, 255}:    {180, .07, .16},
		{246, 248, 255, 255}: {227, 1, .98},
		{168, 70, 160, 255}:  {305, .41, .47},
		{80, 81, 79, 255}:    {90, .01, .31},
	}

	for c0, d := range colors {
		c1 := d.ToRGBA()
		assert.Equal(t, c0, c1)
	}
}
